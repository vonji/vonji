package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/controllers"
	"github.com/vonji/vonji-api/utils"
	"github.com/vonji/vonji-api/services"
)

func handleError(w http.ResponseWriter, error *utils.HttpError) {
	http.Error(w, error.Error, error.Code)
	println(services.Error.InternalError)
}

func checkError(w http.ResponseWriter, error *utils.HttpError) bool {
	if services.Error != nil {
		handleError(w, services.Error)
		services.Error = nil
		return true
	}
	if error != nil {
		handleError(w, error)
		return true
	}
	return false
}

var GetOneHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ParseUint(mux.Vars(r)["id"])
		if err != nil {
			handleError(w, utils.BadRequest("Parameter ID is not an unsigned integer"))
			return
		}
		obj, httpErr := ctrl.GetOne(id)

		if checkError(w, httpErr) {
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}

var GetAllHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := ctrl.GetAll()

		if checkError(w, err) {
			return
		}

		json.NewEncoder(w).Encode(obj)
	}
}

var PostHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := ctrl.Create(w, r)

		if checkError(w, err) {
			return
		}

		json.NewEncoder(w).Encode(obj)
		w.WriteHeader(http.StatusCreated)
	}
}

var PutHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if checkError(w, ctrl.Update(w, r)) {
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

var DeleteHandler = func(ctrl controllers.APIController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctrl := controllers.RequestController{}

		id, err := utils.ParseUint(mux.Vars(r)["id"])
		if err != nil {
			handleError(w, utils.BadRequest("Parameter ID is not an unsigned integer"))
			return
		}

		if checkError(w, ctrl.Delete(id)) {
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

//The order is important
func RegisterRoutes(r *mux.Router) {
	ru := r.PathPrefix("/users").Subrouter()
	ru.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(GetOneHandler(controllers.UserController{}))
	ru.Methods("GET").PathPrefix("/{email}").HandlerFunc(controllers.GetUserByEmail)
	ru.Methods("GET").HandlerFunc(GetAllHandler(controllers.UserController{}))
	ru.Methods("POST").HandlerFunc(PostHandler(controllers.UserController{}))
	ru.Methods("PUT").HandlerFunc(PutHandler(controllers.UserController{}))
	ru.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(DeleteHandler(controllers.UserController{}))

	rq := r.PathPrefix("/requests").Subrouter()
	rq.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(GetOneHandler(controllers.RequestController{}))
	rq.Methods("GET").HandlerFunc(GetAllHandler(controllers.RequestController{}))
	rq.Methods("POST").HandlerFunc(PostHandler(controllers.RequestController{}))
	rq.Methods("PUT").HandlerFunc(PutHandler(controllers.RequestController{}))
	rq.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(DeleteHandler(controllers.RequestController{}))

	rs := r.PathPrefix("/responses").Subrouter()
	rs.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(GetOneHandler(controllers.ResponseController{}))
	rs.Methods("GET").HandlerFunc(GetAllHandler(controllers.ResponseController{}))
	rs.Methods("POST").HandlerFunc(PostHandler(controllers.ResponseController{}))
	rs.Methods("PUT").HandlerFunc(PutHandler(controllers.ResponseController{}))
	rs.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(DeleteHandler(controllers.ResponseController{}))

	rc := r.PathPrefix("/comments").Subrouter()
	rc.Methods("GET").PathPrefix("/{id:[0-9]+}").HandlerFunc(GetOneHandler(controllers.CommentController{}))
	rc.Methods("GET").HandlerFunc(GetAllHandler(controllers.CommentController{}))
	rc.Methods("POST").HandlerFunc(PostHandler(controllers.CommentController{}))
	rc.Methods("PUT").HandlerFunc(PutHandler(controllers.CommentController{}))
	rc.Methods("DELETE").PathPrefix("/{id:[0-9]+}").HandlerFunc(DeleteHandler(controllers.CommentController{}))
}
