package users

import (
	"encoding/json"
	"net/http"

	"github.com/devdinu/simple-api/logger"
)

type credential struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func Authenticate(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req credential
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Errorf("[Authentication] error parsing request: %v", err)
		}
		if !svc.Authenticate(r.Context(), req) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message":"unauthorised user"}`))
			return
		}
	}
}
