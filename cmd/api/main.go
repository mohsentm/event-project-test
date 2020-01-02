package main

import (
	"github.com/gorilla/mux"
	"github.com/mohsentm/event-project-test/config"
	"github.com/mohsentm/event-project-test/internal/router"
	"log"
	"net/http"
	"time"
)

func init() {
	config.Init()
}

func main() {
	conf := config.Get()

	r := mux.NewRouter().StrictSlash(true)
	subRouter := r.PathPrefix("/api").Subrouter()

	srv := &http.Server{
		Handler:      router.Routes(subRouter),
		Addr:         ":" + conf.ApiPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
