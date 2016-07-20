package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type ResponseController struct {
	APIBaseController
}

func (ctrl ResponseController) GetAll() (interface{}, *utils.HttpError) {
	responses := []models.Response{}
	ctrl.GetDB().Find(&responses)

	for i, response := range responses {
		ctrl.GetDB().Model(&response).Related(&responses[i].User)
		ctrl.GetDB().Where(&models.Comment{ ResponseID: response.ID }).Find(&response.Comments)
		for j, comment := range response.Comments {
			ctrl.GetDB().Model(&comment).Related(&response.Comments[j].User)
		}
	}

	return responses, nil
}

func (ctrl ResponseController) GetOne(id uint) (interface{}, *utils.HttpError) {
	response := models.Response{}

	ctrl.GetDB().First(&response, id)
	if err := ctrl.CheckID(response.ID); err != nil {
		return nil, err
	}

	ctrl.GetDB().Model(&response).Related(&response.User)
	ctrl.GetDB().Where(&models.Comment{ ResponseID: response.ID }).Find(&response.Comments)
	for i, comment := range response.Comments {
		ctrl.GetDB().Model(&comment).Related(&response.Comments[i].User)
	}

	return response, nil
}

func (ctrl ResponseController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	response := models.Response{}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctrl.GetDB().Create(&response)

	return response, nil
}

func (ctrl ResponseController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError{
	response := models.Response{}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctrl.GetDB().Save(&response)

	return nil
}

func (ctrl ResponseController) Delete(id uint) *utils.HttpError {
	response := models.Response{}

	response.ID = id

	ctrl.GetDB().Delete(&response)

	return nil
}