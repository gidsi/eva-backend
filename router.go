package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range IndexRoutes {
		var handler http.Handler
		handler = Logger(route.Handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)

		router.
		Methods("OPTIONS").
		Name("Options Handler").
		Handler(http.HandlerFunc(optionsHandler))
	}

	router.
	Methods("OPTIONS").
	Name("Options Handler").
	Handler(http.HandlerFunc(optionsHandler))

	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Print("---------------------------------")
	log.Print(w)
	log.Print("---------------------------------")
	log.Print(r)
	log.Print("---------------------------------")
	w.WriteHeader(http.StatusNotFound)
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.WriteHeader(200)
}