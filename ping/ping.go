package ping

import (
	"net/http"

	"github.com/devdinu/simple-api/logger"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{"response":"pong"}`)); err != nil {
			logger.Errorf("[Ping] error writing response: %v", err)
		}
	}
}
