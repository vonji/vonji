package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
	"github.com/gorilla/mux"
)

type NotificationController struct {
	APIBaseController
}

func (ctrl NotificationController) GetAll() (interface{}, *utils.HttpError) {
	return services.Notification.GetAll(), nil
}

func (ctrl NotificationController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Notification.GetOne(id), nil
}

func (ctrl NotificationController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	notification := models.Notification{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &notification); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Notification.GetOneWhere(&notification), nil
}

func (ctrl NotificationController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	notification := models.Notification{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &notification); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Notification.GetAllWhere(&notification), nil
}

func (ctrl NotificationController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	notification := models.Notification{}

	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Notification.Create(&notification), nil
}

func (ctrl NotificationController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	notification := models.Notification{}

	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Notification.Update(&notification)

	return nil
}

func (ctrl NotificationController) Delete(id uint) *utils.HttpError {
	services.Notification.Delete(id)

	return nil
}
