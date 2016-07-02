package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/controllers"
)

func parseUint(s string) (uint, error) { //TODO move
	n, err := strconv.ParseUint(s, 10, 64)
	return uint(n), err
}

var GetOneHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.RequestController{}
		id, err := parseUint(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(ctrl.GetOne(id))
	}
}

var GetAllHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(ctrl.GetAll())
	}
}

var PostHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ctrl.Create(w, r))
	}
}

var PutHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl.Update(w, r)
		w.WriteHeader(http.StatusNoContent)
	}
}

var DeleteHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.RequestController{}
		id, err := parseUint(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
			return
		}
		ctrl.Delete(id, w, r)
		w.WriteHeader(http.StatusNoContent)
	}
}

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

	rq.Methods("GET").PathPrefix("/{id:[0-9]+}").
		HandlerFunc(GetOneHandler(controllers.RequestController{}))
	rq.Methods("GET").
		HandlerFunc(GetAllHandler(controllers.RequestController{}))
	rq.Methods("POST").
		HandlerFunc(PostHandler(controllers.RequestController{}))
	rq.Methods("PUT").
		HandlerFunc(PutHandler(controllers.RequestController{}))
	rq.Methods("DELETE").PathPrefix("/{id:[0-9]+}").
		HandlerFunc(DeleteHandler(controllers.RequestController{}))

	rs := r.PathPrefix("/responses").Subrouter()
	rs.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.GetResponseById)
	rs.Methods("GET").HandlerFunc(controllers.GetResponse)
	rs.Methods("POST").HandlerFunc(controllers.CreateResponse)
	rs.Methods("PUT").HandlerFunc(controllers.UpdateResponse)
	rs.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(controllers.DeleteResponse)
}
