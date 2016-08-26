package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
	"github.com/gorilla/mux"
)

type RequestController struct {
	APIBaseController
}

func (ctrl RequestController) GetAll() (interface{}, *utils.HttpError) {
	return services.Request.GetAll(), nil
}

func (ctrl RequestController) Light() (interface{}, *utils.HttpError) {
	return services.Request.Light(), nil
}

func (ctrl RequestController) GetOne(id uint) (interface{}, *utils.HttpError) {
	request := services.Request.GetOne(id)

	if services.Error == nil {
		go (func() {
			request.Views++
			ctrl.GetDB().Save(&request)
		})()
	}

	return request, nil
}

func (ctrl RequestController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	request := models.Request{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &request); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Request.GetOneWhere(&request), nil
}

func (ctrl RequestController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	request := models.Request{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &request); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Request.GetAllWhere(&request), nil
}

func (ctrl RequestController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	request := models.Request{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Request.Create(&request), nil
}

func (ctrl RequestController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	request := models.Request{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Request.Update(&request)

	return nil
}

func (ctrl RequestController) Delete(id uint) *utils.HttpError {
	services.Request.Delete(id)
	return nil
}
