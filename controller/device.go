package controller

import (
	"log"
	"net/http"
)

func (ctrl *HttpController) CreateTemperature(w http.ResponseWriter, r *http.Request) {
	log.Println("hitting the post request")
}
