package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Path    string
	Methods []string
	Handler http.Handler
}

type Routes []Route

var routes = Routes{
	Route{
		"/calls/incoming",
		[]string{"GET"},
		NewAppHandler(&IncomingConferenceHandler{}),
	},
	Route{
		"/calls/conferences",
		[]string{"POST"},
		NewAppHandler(&IncomingConferenceHandler{}),
	},
}

// Router defines the routes for the API and
// returns a *mux.Router
func Router() *mux.Router {

	r := mux.NewRouter()

	for _, route := range routes {
		r.Handle(route.Path, route.Handler).Methods(route.Methods...)
	}

	return r
}