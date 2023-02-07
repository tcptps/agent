package jobapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

func AuthMdlw(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(w, r)

			enc := json.NewEncoder(w)

			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.WriteHeader(http.StatusUnauthorized)
				enc.Encode(ErrorResponse{Error: "authorization header is required"})
				return
			}

			authType, reqToken, found := strings.Cut(auth, " ")
			if !found {
				w.WriteHeader(http.StatusUnauthorized)
				enc.Encode(ErrorResponse{Error: "invalid authorization header: must be in the form `Bearer <token>`"})
				return
			}

			if authType != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				enc.Encode(ErrorResponse{Error: "invalid authorization header: type must be Bearer"})
				return
			}

			if reqToken != token {
				w.WriteHeader(http.StatusUnauthorized)
				enc.Encode(ErrorResponse{Error: "invalid authorization token"})
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
