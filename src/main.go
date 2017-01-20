package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"pkg/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/godbtest/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/godbtest/users/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/godbtest/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/godbtest/users/{id}", controller.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":3030", router)
}
