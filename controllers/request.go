package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/models"
	"github.com/vonji/vonji-api/api"
	"github.com/jinzhu/gorm"
)

func GetRequests(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()

	requests := []models.Request{}
	ctx.Db.Find(&requests)

	for i, request := range requests {
		ctx.Db.Model(&request).Related(&requests[i].User)
		ctx.Db.Model(&request).Related(&requests[i].Responses)
		for j, response := range requests[i].Responses {
			ctx.Db.Model(&response).Related(&requests[i].Responses[j].User)
		}
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

	if request.ID == 0 {
		http.Error(w, fmt.Sprintf("No request with ID %d found", id), http.StatusNotFound)
		return
	}

	ctx.Db.Model(&request).Related(&request.User)
	ctx.Db.Model(&request).Related(&request.Responses)
	for i, response := range request.Responses {
		ctx.Db.Model(&response).Related(&request.Responses[i].User)
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
	json.NewEncoder(w).Encode(models.User{ Model: gorm.Model { ID: request.ID } })
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

	ctx.Db.Delete(&request)
	ctx.Db.Where(&models.Response{ RequestID: request.ID }).Delete(&models.Response{})
}
