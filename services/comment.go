package services

import (
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/utils"
)

type CommentService struct {
	BaseService
}

func (service CommentService) GetAll() []models.Comment {
	if Error != nil {
		return nil
	}

	comments := []models.Comment{}
	service.GetDB().Find(&comments)

	for i, comment := range comments {
		comments[i].User = User.GetOne(comment.UserID)
	}

	return comments
}

func (service CommentService) GetOne(id uint) models.Comment {
	if Error != nil {
		return models.Comment{}
	}

	comment := models.Comment{}

	if err := service.GetDB().First(&comment, id); err.Error != nil {
		Error = utils.DatabaseError(err)
		return models.Comment{}
	}

	comment.User = User.GetOne(comment.UserID)

	return comment
}

func (service CommentService) GetOneWhere(comment *models.Comment) models.Comment {
	if Error != nil {
		return models.Comment{}
	}

	if err := service.GetDB().Where(&comment).First(&comment); err.Error != nil {
		Error = utils.DatabaseError(err)
		return models.Comment{}
	}

	comment.User = User.GetOne(comment.UserID)

	return *comment
}