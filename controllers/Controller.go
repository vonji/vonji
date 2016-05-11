package controllers

import (
	"github.com/gorilla/mux"
)

type Controller interface {//todo maybe this interface is useless
	Init(r *mux.Router)
}