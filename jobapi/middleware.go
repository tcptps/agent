package jobapi

import (
	"errors"
	"net/http"
	"strings"
)

func AuthMdlw(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(w, r)

			auth := r.Header.Get("Authorization")
			if auth == "" {
				writeError(w, errors.New("authorization header is required"), http.StatusUnauthorized)
				return
			}

			authType, reqToken, found := strings.Cut(auth, " ")
			if !found {
				writeError(w, errors.New("invalid authorization header: must be in the form `Bearer <token>`"), http.StatusUnauthorized)
				return
			}

			if authType != "Bearer" {
				writeError(w, errors.New("invalid authorization header: type must be Bearer"), http.StatusUnauthorized)
				return
			}

			if reqToken != token {
				writeError(w, errors.New("invalid authorization token"), http.StatusUnauthorized)
				return
			}
		})
	}
}

func HeadersMdlw() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(w, r)

			w.Header().Set("Content-Type", "application/json")
		})
	}
}
