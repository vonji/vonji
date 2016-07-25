package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
)

type CommentController struct {
	APIBaseController
}

func (ctrl CommentController) GetAll() (interface{}, *utils.HttpError) {
	return services.Comment.GetAll(), nil
}

func (ctrl CommentController) GetOne(id uint) (interface{}, *utils.HttpError) {
	return services.Comment.GetOne(id), nil
}

func (ctrl CommentController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	comment := models.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return nil, utils.BadRequest(err.Error())
	}

	return services.Comment.Create(&comment), nil
}

func (ctrl CommentController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	comment := models.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return utils.BadRequest(err.Error())
	}

	services.Comment.Update(&comment)

	return nil
}

func (ctrl CommentController) Delete(id uint) *utils.HttpError {
	services.Comment.Delete(id)

	return nil
}
