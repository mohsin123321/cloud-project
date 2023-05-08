package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/mohsin123321/cloud-project/config"
	"github.com/mohsin123321/cloud-project/controller"
	"github.com/mohsin123321/cloud-project/dataservice"
	_ "github.com/mohsin123321/cloud-project/docs"
	"github.com/mohsin123321/cloud-project/router"
	"github.com/mohsin123321/cloud-project/utility"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
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
	if !reflect.ValueOf(ds.Db).IsZero() {
		ds.Db.Close()
	}

	utility := utility.Utility{}
	ctrl := controller.HttpController{
		Ds: ds,
		Ut: &utility,
	}
	r := mux.NewRouter()

	router.SetupRoutes(r, &ctrl)

	if config.Config.ShowDocs {
		r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	}

	// start server
	log.Println("server is listening on port:", config.Config.Server.Port)

	// Since we don't have to start server with TLS so we ignore this rule
	// nosemgrep
	err := http.ListenAndServe(":"+config.Config.Server.Port, cors.AllowAll().Handler(r))

	if err != nil {
		log.Println(err)
	}
}
