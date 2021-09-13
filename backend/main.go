package main

import (
	"log"
	"os"
	"net/http"

	"github.com/zohaib194/cv-finder/backend/controller"
	"github.com/zohaib194/cv-finder/backend/model"
)

func main() {

	port := os.Getenv("GOPORT");
	filePath := os.Getenv("CVFILEPATH");

	if port == "" {
		log.Fatal("PORT is not set");
	}

	if filePath == "" {
		log.Fatal("CVFILEPATH is not set");
	}

	model.Filepath = filePath;
 
	http.HandleFunc("/api/cv", controller.CVController{}.HandleCVForm)
    log.Println("Listing for requests at http://localhost:" + port + "/api")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}