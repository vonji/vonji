package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/api"
)

func GetRequests(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()

	requests := []models.Request{}
	ctx.Db.Find(&requests)

	for i, request := range requests {
		ctx.Db.Model(&request).Related(&requests[i].User)
	}

	json.NewEncoder(w).Encode(requests)
}

func GetRequestById(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()
	request := models.Request{}

	id, err := parseUint(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
		return
	}

	ctx.Db.First(&request, id)
	ctx.Db.Model(&request).Related(&request.User)

	if request.ID == 0 {
		http.Error(w, fmt.Sprintf("No request with ID %d found", id), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(request)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Db.Create(&request)
}

func UpdateRequest(w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Db.Save(&request)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	request := models.Request{}
	ctx := api.GetContext()

	id, err := parseUint(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
		return
	}

	request.ID = id

	ctx.Db.Delete(&request)//Soft delete
}
