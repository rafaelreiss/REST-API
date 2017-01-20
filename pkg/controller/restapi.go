package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rafaelreiss/SimpleRestApi/pkg/dao"
)

// Returns all the users
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	user, err := dao.GetAllUsers()
	if err != nil {
		handleError(err, "Failed to load db users: %v", w)
		return
	}

	json, err := json.Marshal(user)

	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
	}
	w.Write(json)
}

// Return an user by his id
func GetUserById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	user, err := dao.FindById(id)
	if err != nil {
		handleError(err, "Failed to read: %v", w)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		handleError(err, "Failed to marshal %v", w)
		return
	}
	w.Write(json)
}

// Creates an user
func CreateUser(w http.ResponseWriter, req *http.Request) {

	user := dao.User{}

	json.NewDecoder(req.Body).Decode(&user)

	if err := dao.Create(user); err != nil {
		handleError(err, "Faile to save :%v", w)
		return
	}
	//uj, _ := json.Marshal(user)
	//fmt.Fprintf(w, "%s", uj)

	w.Write([]byte("OK"))
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := dao.Delete(id); err != nil {
		handleError(err, "Failed to remove user: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

//Errors handlers
func handleError(err error, msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(msg, err)))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
