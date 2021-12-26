package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("hitting the get request")
	fmt.Fprintf(w, "hello \n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	log.Println("Starting goBase server")
	r := mux.NewRouter()
	r.HandleFunc("/device", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ciao tutti"))
		log.Println("hitting the post request")
	}).Methods("POST")
	r.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	//port := "8080"
	err := http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	if err != nil {
		log.Println(err)
	}
}
