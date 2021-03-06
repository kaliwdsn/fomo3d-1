package main

import (
	"net/http"

	"github.com/go-ray/fomo3d/controller"
	"github.com/go-ray/logging"
	"github.com/gorilla/mux"
)

func setHandles(r *mux.Router) {
	r.Path("/").HandlerFunc(controller.Root).Methods("GET")
	r.Path("/api/players").Queries("offset", "{offset}").Queries("num", "{amount}").HandlerFunc(controller.PlayersHandler).Methods("GET")

	r.HandleFunc("/api/players/{offset}/{amount}", controller.PlayersHandler).Methods("GET")
	r.Path("/api/players").HandlerFunc(controller.PlayersHandler).Methods("GET")

	r.HandleFunc("/api/names", controller.NamesHandler).Queries("offset", "{offset}", "num", "{amount}").Methods("GET")
	r.HandleFunc("/api/names/{offset}/{amount}", controller.NamesHandler).Methods("GET")
	r.HandleFunc("/api/names", controller.NamesHandler).Methods("GET")
	r.HandleFunc("/api/keyHolderStats", controller.KeyHolderStatsHandler).Queries("address", "{address}").Methods("GET")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logging.Debug("uri:", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
