package main

import (
	"github.com/gorilla/mux"
	"github.com/mohsentm/event-project-test/src/router"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	subRouter := r.PathPrefix("/api").Subrouter()

	log.Fatal(http.ListenAndServe(":8080", router.Routes(subRouter)))
}
