package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
	"github.com/gorilla/mux"
)

type AdController struct {
	APIBaseController
}

func (ctrl AdController) GetAll() (interface{}, *utils.HttpError) {
	return services.Ad.GetAll(), nil
}

func (ctrl AdController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Ad.GetOne(id), nil
}

func (ctrl AdController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	ad := models.Ad{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &ad); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Ad.GetOneWhere(&ad), nil
}

func (ctrl AdController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	ad := models.Ad{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &ad); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Ad.GetAllWhere(&ad), nil
}

func (ctrl AdController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	ad := models.Ad{}

	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Ad.Create(&ad), nil
}

func (ctrl AdController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	ad := models.Ad{}

	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Ad.Update(&ad)

	return nil
}

func (ctrl AdController) Delete(id uint) *utils.HttpError {
	services.Ad.Delete(id)

	return nil
}
