package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
}

func (c *UserController) Init(router *mux.Router) {
	router.HandleFunc("/user", Get).Methods("GET")
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GETUSER\n"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GETUSER\n"))
}

func Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GETUSER\n"))
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GETUSER\n"))
}
