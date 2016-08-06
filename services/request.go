package services

import (
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/models"
)

type RequestService struct {
	BaseService
}

func (service RequestService) GetAll() []models.Request {
	if Error != nil {
		return nil
	}

	requests := []models.Request{}

	if db := service.GetDB().Find(&requests); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	for i, request := range requests {
		requests[i].User = *User.GetOne(request.UserID)
		requests[i].Responses = Response.GetAllWhere(&models.Response{ RequestID: request.ID })
		requests[i].Comments = Comment.GetAllWhere(&models.Comment{ RequestID: request.ID })
		if db := service.GetDB().Model(&request).Association("tags").Find(&requests[i].Tags); db.Error != nil {
			Error = utils.AssociationError(db)
			return nil
		}
	}

	return requests
}

func (service RequestService) GetOne(id uint) *models.Request {
	if Error != nil {
		return nil
	}

	request := models.Request{}

	if db:= service.GetDB().First(&request, id); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	request.User = *User.GetOne(request.UserID)
	request.Responses = Response.GetAllWhere(&models.Response{ RequestID: id })
	request.Comments = Comment.GetAllWhere(&models.Comment{ RequestID: id })
	if db := service.GetDB().Model(&request).Association("tags").Find(&request.Tags); db.Error != nil {
		Error = utils.AssociationError(db)
		return nil
	}

	go (func() {
		request.Views++
		if db := service.GetDB().Save(&request); db.Error != nil {
			println(utils.DatabaseError(db).InternalError)
		}
	})()

	return &request
}

func (service RequestService) Create(request *models.Request) *models.Request {
	if Error != nil {
		return nil
	}

	request.Tags = Tag.GetOrCreate(request.Tags)
	if db := service.GetDB().Create(&request); db.Error != nil {
		Error = utils.DatabaseError(db)
		return nil
	}

	return Request.GetOne(request.ID)
}

func (service RequestService) Update(request *models.Request) {
	if Error != nil {
		return
	}

	request.Tags = Tag.GetOrCreate(request.Tags)
	if db := service.GetDB().Save(&request); db.Error != nil {
		Error = utils.DatabaseError(db)
	}
}

func (service RequestService) Delete(id uint) {
	if Error != nil {
		return
	}

	request := models.Request{}
	request.ID = id

	if db := service.GetDB().Delete(&request); db.Error != nil {
		Error = utils.DatabaseError(db)
		return
	}
	//delete? service.GetDB().Where(&models.Response{RequestID: request.ID}).Delete(&models.Response{})
}

