package controllers

import (
	"github.com/gorilla/mux"
)

//TODO maybe this interface is useless
type Controller interface {
	Init(r *mux.Router)
}
