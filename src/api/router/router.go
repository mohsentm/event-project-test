package router

import (
	"github.com/gorilla/mux"
	"net/http"
	eventHandler "restapi/src/event"
	"restapi/src/home"
)

func Routes(router *mux.Router) *mux.Router {
	router = homeRoutes(router)
	router = eventRoutes(router)
	return router
}

func homeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", home.Handler)
	return router
}

func eventRoutes(router *mux.Router) *mux.Router {
	subRouter := router.PathPrefix("/event").Subrouter()
	subRouter.HandleFunc("/", eventHandler.GetAllEvent).Methods(http.MethodGet)
	subRouter.HandleFunc("/", eventHandler.CreateEvent).Methods(http.MethodPost)
	subRouter.HandleFunc("/{id}", eventHandler.GetOneEvent).Methods(http.MethodGet)
	subRouter.HandleFunc("/{id}", eventHandler.UpdateEvent).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", eventHandler.DeleteEvent).Methods(http.MethodDelete)

	return router
}
