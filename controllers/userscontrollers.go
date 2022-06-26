package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-management/database"
	"user-management/entity"

	"github.com/gorilla/mux"
)

//get all user data
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var user []entity.User
	database.Connector.Find(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//delete specific user by ID
func DeletUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key, vars)

	var user entity.User
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}

//create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)

	database.Connector.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

//update user by ID
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)
	database.Connector.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
