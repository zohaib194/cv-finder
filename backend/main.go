package main

import (
	"log"
	"os"
	"net/http"

	"github.com/zohaib194/cv-finder/backend/controller"
)

func main() {
	// Hello world, the web server

	port := os.Getenv("GOPORT");
	if port == "" {
		log.Fatal("PORT is not set");
	}

	http.HandleFunc("/api/cv", controller.HandleCVForm)
    log.Println("Listing for requests at http://localhost:" + port + "/api")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}