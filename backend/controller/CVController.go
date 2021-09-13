package controller

import (
	"net/http"
	"encoding/json"
	
	"github.com/zohaib194/cv-finder/backend/model"
)

type CVController struct {

}

func (cvController CVController) HandleCVForm(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	http.Header.Add(w.Header(), "Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		
	} else if r.Method == "GET" {
		model.ReadAllCVs();
		cv := model.ReadAllCVs();

		w.WriteHeader(http.StatusCreated)
    	json.NewEncoder(w).Encode(cv);
	}
}