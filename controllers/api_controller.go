package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type APIController interface {
	GetAll() interface{}
	GetOne(id uint) interface{}
	Create(w http.ResponseWriter, r *http.Request) interface{}
	Update(w http.ResponseWriter, r *http.Request)
	Delete(id uint, w http.ResponseWriter, r *http.Request)

	GetDB() *gorm.DB
	GetResponseWriter() http.ResponseWriter
}
