package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohsin123321/cloud-project/controller"
	"github.com/mohsin123321/cloud-project/dataservice"
	"github.com/mohsin123321/cloud-project/router"
	"github.com/rs/cors"
)

func main() {
	log.Println("Starting the server")

	// Initialize the dataservice
	ds := dataservice.SetupDGS()
	// release the resoures by closing connection
	defer ds.Db.Close()

	ctrl := controller.HttpController{
		DS: ds,
	}
	r := mux.NewRouter()
	router.SetupRoutes(r, &ctrl)

	port := "8080"
	err := http.ListenAndServe(":"+port, cors.AllowAll().Handler(r))
	if err != nil {
		log.Println(err)
	}
}
