package controllers

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/vonji/vonji-api/api"
)

type APIBaseController struct {
	ResponseWriter http.ResponseWriter
}

func (ctrl APIBaseController) CheckID(id uint) {
	if id == 0 {
		http.Error(ctrl.GetResponseWriter(), fmt.Sprintf("No request with ID %d found", id), http.StatusNotFound)
	}
}

func (ctrl APIBaseController) GetDB() *gorm.DB {
	return api.GetContext().DB
}

func (ctrl APIBaseController) GetResponseWriter() http.ResponseWriter {
	return ctrl.ResponseWriter
}
