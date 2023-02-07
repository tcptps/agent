package jobapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) router() chi.Router {
	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
		// middleware.Logger, // REVIEW: Should we log requests to this API? If so, where should we log them to? The job logs?
		s.authMdlw,
		s.headersMdlw,
	)

	r.Route("/api/current-job/v0", func(r chi.Router) {
		r.Get("/env", s.getEnv())
		r.Patch("/env", s.patchEnv())
		r.Delete("/env", s.deleteEnv())
	})

	return r
}

func (s *Server) getEnv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// STUB
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) patchEnv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// STUB
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) deleteEnv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// STUB
		w.WriteHeader(http.StatusOK)
	}
}
