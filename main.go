package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/controller"
	"github.com/mohsin123321/cloud-project/dataservice"
	_ "github.com/mohsin123321/cloud-project/docs"
	"github.com/mohsin123321/cloud-project/router"
	"github.com/mohsin123321/cloud-project/utility"
)

// @title IOT Device API
// @version 1.0.0
// @description API To insert IOT data
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Auth-Token
func main() {
	log.Println("Starting the server by user")

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

	// Initialize the router
	r := chi.NewRouter()

	r.Mount("/api", router.SetupRoutes(&ctrl))

	// start server
	log.Println("server is listening on port:", config.Config.Server.Port)

	// Since we don't have to start server with TLS so we ignore this rule
	// nosemgrep
	err := http.ListenAndServe(":"+config.Config.Server.Port, r)

	if err != nil {
		log.Println(err)
	}
}
