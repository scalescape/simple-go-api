package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/devdinu/simple-api/logger"
)

type Users struct {
	svc Service
}

type response struct {
	Count string    `db:"count"`
	Time  time.Time `db:"time"`
}

func CountUsersHandler(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cnt, err := svc.TotalUsers(r.Context())
		if err != nil {
			writeError(w, http.StatusInternalServerError, "list users", err)
			return
		}
		resp := response{Count: strconv.Itoa(cnt), Time: time.Now()}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("[Users.List] error writing response: %v", err)
		}
	}
}

func ListUsersHandler(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := svc.ListUsers(r.Context())
		if err != nil {
			writeError(w, http.StatusInternalServerError, "list users", err)
			return
		}
		if err := json.NewEncoder(w).Encode(users); err != nil {
			logger.Errorf("[Users.List] error writing response: %v", err)
		}
	}
}

func writeError(w http.ResponseWriter, status int, message string, err error) {
	resp := struct{ Error string }{Error: message}
	w.WriteHeader(status)
	logger.Errorf("[Users.List] failure : %s with error: %v", message, err)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("[Ping] error writing response: %v", err)
	}
}
