package services

import (
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/models"
	"github.com/jinzhu/gorm"
)

type ResponseService struct {
	BaseService
}

func (service ResponseService) GetAll() []models.Response {
	if Error != nil {
		return nil
	}

	responses := []models.Response{}

	if err := service.GetDB().Find(&responses); err != nil {
		Error = utils.DatabaseError(err)
		return nil
	}

	for i, response := range responses {
		responses[i].User = *User.GetOne(response.UserID)
		responses[i].Comments = Comment.GetAllWhere(&models.Comment{ ResponseID: response.ID })
	}

	return responses
}

func (service ResponseService) GetAllWhere(response *models.Response) []models.Response {
	if Error != nil {
		return nil
	}

	responses := []models.Response{}

	if db := service.GetDB().Where(response).Find(&responses); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	for i, response := range responses {
		responses[i].User = *User.GetOne(response.UserID)
		responses[i].Comments = Comment.GetAllWhere(&models.Comment{ ResponseID: response.ID })
	}

	return responses
}

func (service ResponseService) GetOne(id uint) *models.Response {
	if Error != nil {
		return nil
	}

	return Response.GetOneWhere(&models.Response{ Model: gorm.Model { ID: id } })
}

func (service ResponseService) GetOneWhere(response *models.Response) *models.Response {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Where(response).First(&response); db.Error != nil {
		if !db.RecordNotFound() {
			Error = utils.DatabaseError(db)
		}
		return nil
	}

	response.User = *User.GetOne(response.UserID)
	response.Comments = Comment.GetAllWhere(&models.Comment{ ResponseID: response.ID })

	return response
}

func (service ResponseService) Create(response *models.Response)  *models.Response {
	if Error != nil {
		return nil
	}

	if db := service.GetDB().Create(&response); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	go(func() {
		request := Request.GetOne(response.RequestID)
		Notification.Create(&models.Notification{
			UserID: request.UserID,
			Title: "Votre demande à récue une nouvelle réponse",
			Message: request.Title,
		})
	})()

	return Response.GetOne(response.ID)
}

func (service ResponseService) Update(response *models.Response) {
	if Error != nil {
		return
	}

	if db := service.GetDB().Save(&response); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service ResponseService) Delete(id uint) {
	if Error != nil {
		return
	}

	response := models.Response{}
	response.ID = id

	if db := service.GetDB().Delete(&response); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}
