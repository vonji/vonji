package routes

import (
	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/controllers"
)

//TODO auto register default routes (send type in param or something)
//The order is important
func RegisterRoutes(r *mux.Router) {
	ru := r.PathPrefix("/users").Subrouter()
	ru.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.GetUserById)
	ru.Methods("GET").PathPrefix("/{email}").HandlerFunc(controllers.GetUserByEmail)
	ru.Methods("GET").HandlerFunc(controllers.GetUsers)
	ru.Methods("POST").HandlerFunc(controllers.CreateUser)
	ru.Methods("PUT").HandlerFunc(controllers.UpdateUser)
	ru.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.DeleteUser)

	rq := r.PathPrefix("/requests").Subrouter()
	rq.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.GetRequestById)
	rq.Methods("GET").HandlerFunc(controllers.GetRequests)
	rq.Methods("POST").HandlerFunc(controllers.CreateRequest)
	rq.Methods("PUT").HandlerFunc(controllers.UpdateRequest)
	rq.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.DeleteRequest)

	rs := r.PathPrefix("/responses").Subrouter()
	rs.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.GetResponseById)
	rs.Methods("GET").HandlerFunc(controllers.GetResponse)
	rs.Methods("POST").HandlerFunc(controllers.CreateResponse)
	rs.Methods("PUT").HandlerFunc(controllers.UpdateResponse)
	rs.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.DeleteResponse)
}
