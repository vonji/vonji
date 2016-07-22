package utils

import (
	"net/http"
	"fmt"

	"github.com/jinzhu/gorm"
)

//TODO MOVE TO ERROR PACKAGE

type HttpError struct {
	Error string
	Code int
	InternalError string
}

func AssociationError(db *gorm.Association) *HttpError {
	return InternalServerError(db.Error.Error())
}

func DatabaseError(db *gorm.DB) *HttpError {
	if db.RecordNotFound() {
		return &HttpError { http.StatusText(http.StatusNotFound), http.StatusNotFound, db.Error.Error() }
	}
	return InternalServerError(db.Error.Error())
}

func BadRequest(error string) *HttpError {
	return &HttpError { http.StatusText(http.StatusBadRequest), http.StatusBadRequest, error }
}

func InternalServerError(error string) *HttpError {
	return &HttpError { http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, error }
}

func NotFound(id uint) *HttpError {
	return &HttpError { http.StatusText(http.StatusNotFound), http.StatusNotFound, fmt.Sprintf("No object with id: %d was found", id) }
}