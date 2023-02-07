package jobapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) authMdlw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		enc := json.NewEncoder(w)

		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			enc.Encode(ErrorResponse{Error: "authorization header is required"})
			return
		}

		authType, token, found := strings.Cut(auth, " ")
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

		if token != s.token {
			w.WriteHeader(http.StatusUnauthorized)
			enc.Encode(ErrorResponse{Error: "invalid authorization token"})
			return
		}
	})
}

func (s *Server) headersMdlw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		w.Header().Set("Content-Type", "application/json")
	})
}
