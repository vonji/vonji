package services

import (
	"github.com/jinzhu/gorm"
	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/utils"
)

type BaseService struct {
}

var Error *utils.HttpError
var User UserService = UserService{}
var Comment CommentService = CommentService{}

func (service BaseService) GetDB() *gorm.DB {
	return api.GetContext().DB
}