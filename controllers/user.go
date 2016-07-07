package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type UserController struct {
	APIBaseController
}

func (ctrl UserController) GetAll() (interface{}, *utils.HttpError) {
	users := []models.User{}
	ctrl.GetDB().Find(&users)

	for i, user := range users {
		ctrl.GetDB().Model(&user).Association("tags").Find(&users[i].Tags)
	}

	return users, nil
}

func (ctrl UserController) GetOne(id uint) (interface{}, *utils.HttpError) {
	user := models.User{}

	ctrl.GetDB().First(&user, id)

	if err := ctrl.CheckID(user.ID); err != nil {
		return nil, err
	}

	ctrl.GetDB().Model(&user).Association("tags").Find(&user.Tags)

	return user, nil
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()
	user := models.User{}

	user.Email = mux.Vars(r)["email"]

	ctx.DB.Where(&user).First(&user)

	if user.ID == 0 {
		http.Error(w, fmt.Sprintf("No user with email %s was found", user.Email), http.StatusNotFound)
		return
	}

	ctx.DB.Model(&user).Association("tags").Find(&user.Tags)

	json.NewEncoder(w).Encode(user)
}

func (ctrl UserController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctrl.GetDB().Create(&user)

	return user, nil
}

func (ctrl UserController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}
	ctrl.GetDB().Save(&user)

	return nil
}

func (ctrl UserController) Delete(id uint) *utils.HttpError {
	user := models.User{}

	user.ID = id

	ctrl.GetDB().Delete(&user)
	return nil
}
