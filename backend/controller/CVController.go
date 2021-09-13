package controller

import (
	"io"
	"net/http"
)

type CVController struct {

}

func (cvController CVController) HandleCVForm(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	http.Header.Add(w.Header(), "Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		
	}
}