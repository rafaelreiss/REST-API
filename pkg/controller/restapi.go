package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rafaelreiss/SimpleRestApi/pkg/dao"
	"net/http"
)

// Returns all the users
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	user, err := dao.GetAllUsers()
	if err != nil {
		handleError(err, "Failed to load db users: %v", w)
		return
	}

	js, err := json.Marshal(user)

	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
	}
	w.Write(js)
}

// Update a User by his id
func Update(w http.ResponseWriter, req *http.Request) {
	user := dao.User{}

	if err:= json.NewDecoder(req.Body).Decode(&user);err != nil{
		handleError(err, "Error decoding json: %v", w)
		return
	}

	if err := dao.Update(user); err != nil{
		handleError(err, "Failed updating user: %v", w)
		return
	}

	w.Write([]byte("OK"))

}

// Return an user by his id
func GetUserById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	user, err := dao.FindById(id)
	if err != nil {
		handleError(err, "Failed to find user : %v", w)
		return
	}

	js, err := json.Marshal(user)
	if err != nil {
		handleError(err, "Failed to marshal %v", w)
		return
	}
	w.Write(js)
}

// Creates an user
func CreateUser(w http.ResponseWriter, req *http.Request) {

	user := dao.User{}

	json.NewDecoder(req.Body).Decode(&user)

	if err := dao.Create(user); err != nil {
		handleError(err, "Faile to save :%v", w)
		return
	}

	w.Write([]byte("OK"))
}

// Deletes an user
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := dao.Delete(id); err != nil {
		handleError(err, "Failed to remove user: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

// This function returns from the rest api the error message
func handleError(err error, msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(msg, err)))
}