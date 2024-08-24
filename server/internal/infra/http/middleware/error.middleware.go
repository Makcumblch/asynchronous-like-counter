package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Makcumblch/asynchronous-like-counter/internal/app"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func ErrorMW(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			var appError *app.Error
			if errors.As(err, &appError) {
				errResponse, err := json.Marshal(appError)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				w.Write(errResponse)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
