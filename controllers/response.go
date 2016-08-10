package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
	"github.com/gorilla/mux"
)

type ResponseController struct {
	APIBaseController
}

func (ctrl ResponseController) GetAll() (interface{}, *utils.HttpError) {
	return services.Response.GetAll(), nil
}

func (ctrl ResponseController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Response.GetOne(id), nil
}

func (ctrl ResponseController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	response := models.Response{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &response); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Response.GetOneWhere(&response), nil
}

func (ctrl ResponseController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	response := models.Response{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &response); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Response.GetAllWhere(&response), nil
}

func (ctrl ResponseController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	response := models.Response{}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Response.Create(&response), nil
}

func (ctrl ResponseController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError{
	response := models.Response{}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Response.Update(&response)

	return nil
}

func (ctrl ResponseController) Delete(id uint) *utils.HttpError {
	services.Response.Delete(id)
	return nil
}