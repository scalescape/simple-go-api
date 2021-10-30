package main

import (
	"net/http"

	"github.com/devdinu/simple-api/config"
	"github.com/devdinu/simple-api/ping"
	"github.com/gorilla/mux"
)

func server(appCfg config.Application) (*mux.Router, error) {
	m := mux.NewRouter()
	m.Use(mux.MiddlewareFunc(contentWriter))
	m.Use(mux.CORSMethodMiddleware(m))

	m.HandleFunc("/ping", ping.Handler()).Methods(http.MethodGet, http.MethodOptions)
	return m, nil
}

func contentWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
