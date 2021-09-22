package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zohaib194/cv-finder/backend/controller"
)

func main() {

	port := os.Getenv("GOPORT");
	frontendPath := os.Getenv("FRONTENDPATH");
	router := mux.NewRouter();

	if port == "" {
		log.Fatal("PORT is not set");
	}	
	
	if frontendPath == "" {
		log.Fatal("FRONTENDPATH is not set");
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%s",path);

	router.HandleFunc("/api/v1/cv", controller.CVController{}.HandleCVForm).Methods("GET", "POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist/webapp/")))
    log.Println("Listing for requests at http://localhost:" + port + "/" + "\n");
	log.Fatal(http.ListenAndServe(":"+port, router))

}