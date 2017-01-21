package main

import (
	"github.com/gorilla/mux"
	"github.com/rafaelreiss/SimpleRestApi/pkg/controller"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/godbtest/users", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/godbtest/users/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/godbtest/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/godbtest/users/{id}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/godbtest/users/{id}", controller.Update).Methods("PUT")
	http.ListenAndServe(":3030", router)
}
