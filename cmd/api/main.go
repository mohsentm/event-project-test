package main

import (
	"github.com/gorilla/mux"
	"github.com/mohsentm/event-project-test/internal/router"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	subRouter := r.PathPrefix("/api").Subrouter()

	srv := &http.Server{
		Handler:      router.Routes(subRouter),
		Addr:         ":80",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
