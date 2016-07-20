package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type CommentController struct {
	APIBaseController
}

func (ctrl CommentController) GetAll() (interface{}, *utils.HttpError) {
	comments := []models.Comment{}
	ctrl.GetDB().Find(&comments)

	for i, comment := range comments {
		ctrl.GetDB().Model(&comment).Related(&comments[i].User)
	}

	return comments, nil
}

func (ctrl CommentController) GetOne(id uint) (interface{}, *utils.HttpError) {
	comment := models.Comment{}

	ctrl.GetDB().First(&comment, id)
	if err := ctrl.CheckID(comment.ID); err != nil {
		return nil, err
	}

	ctrl.GetDB().Model(&comment).Related(&comment.User)

	return comment, nil
}

func (ctrl CommentController) Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError) {
	comment := models.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return nil, &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctrl.GetDB().Create(&comment)

	return comment, nil
}

func (ctrl CommentController) Update(w http.ResponseWriter, r *http.Request) *utils.HttpError {
	comment := models.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return &utils.HttpError{ err.Error(), http.StatusBadRequest }
	}

	ctrl.GetDB().Save(&comment)

	return nil
}

func (ctrl CommentController) Delete(id uint) *utils.HttpError {
	comment := models.Comment{}

	comment.ID = id

	ctrl.GetDB().Delete(&comment)
	ctrl.GetDB().Where(&models.Response{RequestID: comment.ID}).Delete(&models.Response{})

	return nil
}
