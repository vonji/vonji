package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/models"
)

func GetResponse(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()

	responses := []models.Response{}
	ctx.Db.Find(&responses)

	for i, request := range responses {
		ctx.Db.Model(&request).Related(&responses[i].User)
	}

	json.NewEncoder(w).Encode(responses)
}

func GetResponseById(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()
	response := models.Response{}

	id, err := parseUint(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
		return
	}

	ctx.Db.First(&response, id)

	if response.ID == 0 {
		http.Error(w, fmt.Sprintf("No request with ID %d found", id), http.StatusNotFound)
		return
	}

	ctx.Db.Model(&response).Related(&response.User)

	json.NewEncoder(w).Encode(response)
}

func CreateResponse(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Db.Create(&response)

	w.WriteHeader(201)
	w.Write([]byte(strconv.Itoa(int(response.ID))))
}

func UpdateResponse(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx.Db.Save(&response)
}

func DeleteResponse(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	ctx := api.GetContext()

	id, err := parseUint(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
		return
	}

	response.ID = id

	ctx.Db.Delete(&response)
}
