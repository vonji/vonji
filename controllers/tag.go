package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/services"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/models"
	"github.com/gorilla/mux"
)

type TagController struct {
	APIBaseController
}

func (ctrl TagController) GetAll() (interface{}, *utils.HttpError) {
	return services.Tag.GetAll(), nil
}

func (ctrl TagController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Tag.GetOne(id), nil
}

func (ctrl TagController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	tag := models.Tag{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &tag); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Tag.GetOneWhere(&tag), nil
}

func (ctrl TagController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	tag := models.Tag{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &tag); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Tag.GetAllWhere(&tag), nil
}

func (ctrl TagController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	tag := models.Tag{}

	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Tag.Create(&tag), nil
}

func (ctrl TagController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	tag := models.Tag{}

	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Tag.Update(&tag)

	return nil
}

func (ctrl TagController) Delete(id uint) *utils.HttpError {
	services.Tag.Delete(id)
	return nil
}

