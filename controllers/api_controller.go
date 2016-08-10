package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/vonji/vonji-api/utils"
)

type APIController interface {
	GetAll() (interface{}, *utils.HttpError)
	GetAllWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError)
	GetOne(id uint) (interface{}, *utils.HttpError)
	GetOneWhere(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError)
	Create(w http.ResponseWriter, r *http.Request) (interface{}, *utils.HttpError)
	Update(w http.ResponseWriter, r *http.Request) *utils.HttpError
	Delete(id uint) *utils.HttpError

	GetDB() *gorm.DB
	GetResponseWriter() http.ResponseWriter
}
