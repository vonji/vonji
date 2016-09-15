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

	if db:= service.GetDB().Find(&comments); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, comment := range comments {
		comments[i].User = *User.GetOne(comment.UserID)
	}

	return comments
}

func (service CommentService) GetOne(id uint) *models.Comment {
	if Error != nil {
		return nil
	}

	comment := models.Comment{}

	if db := service.GetDB().First(&comment, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	comment.User = *User.GetOne(comment.UserID)

	return &comment
}

func (service CommentService) GetOneWhere(comment *models.Comment) *models.Comment {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Where(&comment).First(&comment); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	comment.User = *User.GetOne(comment.UserID)

	return comment
}

func (service CommentService) GetAllWhere(comment *models.Comment) []models.Comment {
	if Error != nil {
		return nil
	}

	comments := []models.Comment{}

	if db := service.GetDB().Where(comment).Find(&comments); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	for i := range comments {
		comments[i].User = *User.GetOne(comments[i].UserID)
	}

	return comments
}

func (service CommentService) Create(comment *models.Comment) *models.Comment {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&comment); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Comment.GetOne(comment.ID)
}

func (service CommentService) Update(comment *models.Comment) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Save(&comment); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service CommentService) Delete(id uint) {
	if Error != nil {
		return
	}

	comment := models.Comment{}
	comment.ID = id

	if db := service.GetDB().Delete(&comment); db.Error != nil {
		Error = utils.DatabaseError(db)
		return
	}
}