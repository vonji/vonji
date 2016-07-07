package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type RequestController struct {
	APIBaseController
}

func (ctrl RequestController) GetAll() (interface{}, *utils.HttpError) {
	requests := []models.Request{}
	ctrl.GetDB().Find(&requests)

	for i, request := range requests {
		ctrl.GetDB().Model(&request).Related(&requests[i].User)
		ctrl.GetDB().Model(&request).Related(&requests[i].Responses)
		for j, response := range requests[i].Responses {
			ctrl.GetDB().Model(&response).Related(&requests[i].Responses[j].User)
		}
	}

	return requests, nil
}

func (ctrl RequestController) GetOne(id uint) (interface{}, *utils.HttpError) {
	request := models.Request{}
	ctrl.GetDB().First(&request, id)

	if err := ctrl.CheckID(request.ID); err != nil {
		return nil, err
	}

	ctrl.GetDB().Model(&request).Related(&request.User)
	ctrl.GetDB().Model(&request).Related(&request.Responses)
	for i, response := range request.Responses {
		ctrl.GetDB().Model(&response).Related(&request.Responses[i].User)
	}

	return request, nil
}

func (ctrl RequestController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	request := models.Request{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctrl.GetDB().Create(&request)

	return request, nil
}

func (ctrl RequestController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	request := models.Request{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctx.DB.Save(&request)

	return nil
}

func (ctrl RequestController) Delete(id uint) *utils.HttpError {
	request := models.Request{}
	ctx := api.GetContext()

	request.ID = id

	ctx.DB.Delete(&request)
	ctx.DB.Where(&models.Response{RequestID: request.ID}).Delete(&models.Response{})

	return nil
}
