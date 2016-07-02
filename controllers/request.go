package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/models"
)

type RequestController struct {
	APIBaseController
}

func (ctrl RequestController) GetAll() interface{} {
	requests := []models.Request{}
	ctrl.GetDB().Find(&requests)

	for i, request := range requests {
		ctrl.GetDB().Model(&request).Related(&requests[i].User)
		ctrl.GetDB().Model(&request).Related(&requests[i].Responses)
		for j, response := range requests[i].Responses {
			ctrl.GetDB().Model(&response).Related(&requests[i].Responses[j].User)
		}
	}

	return requests
}

func (ctrl RequestController) GetOne(id uint) interface{} {
	request := models.Request{}
	ctrl.GetDB().First(&request, id)

	ctrl.CheckID(id)

	ctrl.GetDB().Model(&request).Related(&request.User)
	ctrl.GetDB().Model(&request).Related(&request.Responses)
	for i, response := range request.Responses {
		ctrl.GetDB().Model(&response).Related(&request.Responses[i].User)
	}

	return request
}

func (ctrl RequestController) Create(w http.ResponseWriter, r *http.Request) interface{} {
	request := models.Request{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	ctx.DB.Create(&request)

	return request
}

func (ctrl RequestController) Update(w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.DB.Save(&request)
}

func (ctrl RequestController) Delete(id uint, w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	ctx := api.GetContext()

	request.ID = id

	ctx.DB.Delete(&request)
	ctx.DB.Where(&models.Response{RequestID: request.ID}).Delete(&models.Response{})
}
