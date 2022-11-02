package main

import (
	"back/pkg/db"
	"back/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/user", h.GetAllUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", h.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", h.DeleteUser).Methods(http.MethodDelete)

	log.Println("REST Api is running")
	http.ListenAndServe(":4000", router)
}
