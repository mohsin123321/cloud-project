package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohsin123321/cloud-project/controller"
)

// routes the request to the right controller

func SetupRoutes(router *mux.Router, ctrl controller.ControllerInterface) {

	router.Use(loggingMiddleware)

	//router.Use(recoveryPanicMdlw)

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	router.HandleFunc("/device", ctrl.InsertData).Methods("POST")
}
