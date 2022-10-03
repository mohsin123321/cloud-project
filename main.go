package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/controller"
	"github.com/mohsin123321/cloud-project/dataservice"
	"github.com/mohsin123321/cloud-project/router"
	"github.com/mohsin123321/cloud-project/utility"
	"github.com/rs/cors"
)

func main() {
	log.Println("Starting the server")

	// read config.json
	config.ReadConfig()

	// Initialize the dataservice
	ds := dataservice.SetupDGS()
	// release the resoures by closing connection
	defer ds.Db.Close()

	utility := utility.Utility{}
	ctrl := controller.HttpController{
		Ds: ds,
		Ut: &utility,
	}
	r := mux.NewRouter()

	router.SetupRoutes(r, &ctrl)

	// Since we don't have to start server with TLS so we ignore this rule
	// nosemgrep
	err := http.ListenAndServe(":"+config.Config.Server.Port, cors.AllowAll().Handler(r))

	if err != nil {
		log.Println(err)
	}

	// start server
	log.Println("server is listening on port:", config.Config.Server.Port)
}
