package ping

import (
	"encoding/json"
	"net/http"

	"github.com/devdinu/simple-api/logger"
	"github.com/jmoiron/sqlx"
)

type response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func Handler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			writeError(w, http.StatusInternalServerError, "failed to connect with db", err)
		}
		resp := response{Message: "pong"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Errorf("[Ping] error writing response: %v", err)
		}
	}
}

func writeError(w http.ResponseWriter, status int, message string, err error) {
	logger.Errorf("[Ping] error pinging: %v", err)
	resp := response{Error: "failed to connect with db"}
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Errorf("[Ping] error writing response: %v", err)
	}
}
