package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vonji/vonji-api/api"
	"github.com/vonji/vonji-api/models"
)

//TODO status code + all responses should be JSON

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()

	users := []models.User{}
	ctx.DB.Find(&users)
	for i, user := range users { //TODO There must be another way to do this
		ctx.DB.Model(&user).Association("tags").Find(&users[i].Tags)
	}

	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()
	user := models.User{}

	id, err := parseUint(mux.Vars(r)["id"]) //TODO find shorter syntax

	if err != nil {
		http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
		return
	}

	ctx.DB.First(&user, id)

	if user.ID == 0 {
		http.Error(w, fmt.Sprintf("No user with ID %d found", id), http.StatusNotFound)
		return
	}

	ctx.DB.Model(&user).Association("tags").Find(&user.Tags)

	json.NewEncoder(w).Encode(user)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := api.GetContext()
	user := models.User{}

	user.Email = mux.Vars(r)["email"]

	ctx.DB.Where(&user).First(&user)

	if user.ID == 0 {
		http.Error(w, fmt.Sprintf("No user with email %s was found", user.Email), http.StatusNotFound)
		return
	}

	ctx.DB.Model(&user).Association("tags").Find(&user.Tags)

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx.DB.Create(&user) //TODO check security
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	ctx := api.GetContext()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx.DB.Save(&user) //TODO check security
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	ctx := api.GetContext()

	id, err := parseUint(mux.Vars(r)["id"]) //TODO find shorter syntax

	if err != nil {
		http.Error(w, "Parameter ID is not an unsigned integer", http.StatusBadRequest)
		return
	}

	user.ID = id

	ctx.DB.Delete(&user) //Soft delete
	//TODO return error if the id does not exist
}
