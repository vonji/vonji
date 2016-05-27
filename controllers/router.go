package controllers

import (
	"github.com/gorilla/mux"
)

//TODO move this file somewhere else
//TODO auto register default routes (send type in param or something)
//The order is important
func RegisterRoutes(r *mux.Router) {
	ru := r.PathPrefix("/users").Subrouter()
	ru.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(GetUserById)
	ru.Methods("GET").HandlerFunc(GetUser)
	ru.Methods("POST").HandlerFunc(CreateUser)
	ru.Methods("PUT").HandlerFunc(UpdateUser)
	ru.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(DeleteUser)
}