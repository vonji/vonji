package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
)

type UserController struct {
	APIBaseController
}

func (ctrl UserController) GetAll() (interface{}, *utils.HttpError) {
	return services.User.GetAll(), nil
}

func (ctrl UserController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.User.GetOne(id), nil
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	user := services.User.GetOneByEmail(mux.Vars(r)["email"])

	if (services.Error != nil) {
		http.Error(w, services.Error.Error, services.Error.Code)
		services.Error = nil
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (ctrl UserController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.User.Create(user), nil
}

func (ctrl UserController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.User.Update(user)

	return nil
}

func (ctrl UserController) Delete(id uint) *utils.HttpError {
	services.User.Delete(id)
	return nil
}
