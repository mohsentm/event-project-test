package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restapi/src/api/router"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	subRouter := r.PathPrefix("/api").Subrouter()

	log.Fatal(http.ListenAndServe(":8080", router.Routes(subRouter)))
}
