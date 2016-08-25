package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
	"github.com/gorilla/mux"
)

type AchievementController struct {
	APIBaseController
}

func (ctrl AchievementController) GetAll() (interface{}, *utils.HttpError) {
	return services.Achievement.GetAll(), nil
}

func (ctrl AchievementController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Achievement.GetOne(id), nil
}

func (ctrl AchievementController) GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	achievement := models.Achievement{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &achievement); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return *services.Achievement.GetOneWhere(&achievement), nil
}

func (ctrl AchievementController) GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	achievement := models.Achievement{}

	if err := json.Unmarshal([]byte(mux.Vars(r)["condition"]), &achievement); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Achievement.GetAllWhere(&achievement), nil
}

func (ctrl AchievementController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	achievement := models.Achievement{}

	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Achievement.Create(&achievement), nil
}

func (ctrl AchievementController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	achievement := models.Achievement{}

	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Achievement.Update(&achievement)

	return nil
}

func (ctrl AchievementController) Delete(id uint) *utils.HttpError {
	services.Achievement.Delete(id)

	return nil
}
