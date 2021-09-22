package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zohaib194/cv-finder/backend/controller"
)

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/api/v1/cv", controller.CVController{}.HandleCVForm).Methods("GET", "POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist/webapp/")))
	log.Fatal(http.ListenAndServe(":8000", router))
}