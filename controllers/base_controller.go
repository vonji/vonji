package controllers

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/utils"
)

type APIBaseController struct {
	ResponseWriter http.ResponseWriter
}

func (ctrl APIBaseController) CheckID(id uint) *utils.HttpError {
	if id == 0 {
		return &utils.HttpError{ fmt.Sprintf("No request with ID %d found", id), http.StatusNotFound }
	}
	return nil
}

func (ctrl APIBaseController) GetDB() *gorm.DB {
	return api.GetContext().DB
}

func (ctrl APIBaseController) GetResponseWriter() http.ResponseWriter {
	return ctrl.ResponseWriter
}
