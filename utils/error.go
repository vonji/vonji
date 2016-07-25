package utils

import (
	"net/http"
	"runtime/debug"

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
		return NotFound(db.Error.Error())
	}
	return InternalServerError(db.Error.Error())
}

func BadRequest(error string) *HttpError {
	debug.PrintStack()
	return &HttpError { http.StatusText(http.StatusBadRequest), http.StatusBadRequest, error }
}

func InternalServerError(error string) *HttpError {
	debug.PrintStack()
	return &HttpError { http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, error }
}

func NotFound(error string) *HttpError {
	debug.PrintStack()
	return &HttpError { http.StatusText(http.StatusNotFound), http.StatusNotFound, error }
}