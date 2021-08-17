package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/getfarook/crud-http-server/api"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", api.HomePage)
	myRouter.HandleFunc("/partners", api.GetAllPartners).Methods("GET")
	myRouter.HandleFunc("/partner/{id}", api.GetPartner).Methods("GET")
	myRouter.HandleFunc("/partner", api.AddParner).Methods("POST")
	myRouter.HandleFunc("/partner/{id}", api.DeleteParner).Methods("DELETE")
	myRouter.HandleFunc("/partner/{id}", api.UpdateParner).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
