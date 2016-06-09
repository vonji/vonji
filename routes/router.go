package routes

import (
	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/controllers"
)

//TODO move this file somewhere else
//TODO auto register default routes (send type in param or something)
//The order is important
func RegisterRoutes(r *mux.Router) {
	ru := r.PathPrefix("/users").Subrouter()
	ru.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.GetUserById)
	ru.Methods("GET").HandlerFunc(controllers.GetUser)
	ru.Methods("POST").HandlerFunc(controllers.CreateUser)
	ru.Methods("PUT").HandlerFunc(controllers.UpdateUser)
	ru.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.DeleteUser)
}
