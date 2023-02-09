package jobapi

import (
	"errors"
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware that checks the Authorization header of an incoming request for a Bearer token
// and checks that that token is the correct one.
func AuthMiddleware(token string) func(http.Handler) http.Handler {
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

// HeadersMiddleware is a middleware that sets the common headers for all responses. At the moment, this is just
// Content-Type: application/json.
func HeadersMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(w, r)

			w.Header().Set("Content-Type", "application/json")
		})
	}
}
