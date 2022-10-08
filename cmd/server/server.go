package main

import (
	"net/http"

	"github.com/devdinu/simple-api/config"
	"github.com/devdinu/simple-api/dbi"
	"github.com/devdinu/simple-api/ping"
	"github.com/devdinu/simple-api/users"
	"github.com/gorilla/mux"
)

func server(appCfg config.Application) (*mux.Router, error) {
	m := mux.NewRouter()
	m.Use(mux.MiddlewareFunc(contentWriter))
	m.Use(mux.CORSMethodMiddleware(m))

	db, err := dbi.NewDB(appCfg.DB)
	if err != nil {
		return nil, err
	}
	usersService := users.NewService(db)
	m.HandleFunc("/ping", ping.Handler(db)).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/users/count", users.CountUsersHandler(usersService)).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/users", users.ListUsersHandler(usersService)).Methods(http.MethodGet, http.MethodOptions)
	m.HandleFunc("/auth", users.Authenticate(usersService)).Methods(http.MethodPost, http.MethodOptions)
	return m, nil
}

func contentWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
