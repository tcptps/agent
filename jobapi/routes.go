package jobapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/buildkite/agent/v3/agent"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/maps"
)

// router returns a chi router with the jobapi routes and appropriate middlewares mounted
func (s *Server) router() chi.Router {
	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
		// middleware.Logger, // REVIEW: Should we log requests to this API? If so, where should we log them to? The job logs?
		HeadersMiddleware(),
		AuthMiddleware(s.token),
	)

	r.Route("/api/current-job/v0", func(r chi.Router) {
		r.Get("/env", s.getEnv)
		r.Patch("/env", s.patchEnv)
		r.Delete("/env", s.deleteEnv)
	})

	return r
}

func (s *Server) getEnv(w http.ResponseWriter, _ *http.Request) {
	normalizedEnv := s.environ.Dump()
	resp := EnvGetResponse{Env: normalizedEnv}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) patchEnv(w http.ResponseWriter, r *http.Request) {
	var req EnvUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		err = fmt.Errorf("failed to decode request body: %w", err)
		writeError(w, err, http.StatusBadRequest)
		return
	}

	added := make([]string, len(req.Env))
	updated := make([]string, len(req.Env))
	protected := checkProtected(maps.Keys(req.Env))

	if len(protected) > 0 {
		err = fmt.Errorf("the following environment variables are protected, and cannot be modified: % v", protected)
		writeError(w, err, http.StatusUnprocessableEntity)
		return
	}

	nils := make([]string, 0, len(req.Env))

	for k, v := range req.Env {
		if v == nil {
			nils = append(nils, k)
		}
	}

	if len(nils) > 0 {
		err = fmt.Errorf("removing environment variables (ie setting them to null) is not permitted on this endpiont. The following keys were set to null: % v", nils)
		writeError(w, err, http.StatusUnprocessableEntity)
		return
	}

	for k, v := range req.Env {
		if _, ok := s.environ.Get(k); ok {
			updated = append(updated, k)
		} else {
			added = append(added, k)
		}
		s.environ.Set(k, *v)
	}

	resp := EnvUpdateResponse{
		Added:   added,
		Updated: updated,
	}

	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) deleteEnv(w http.ResponseWriter, r *http.Request) {
	var req EnvDeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		err = fmt.Errorf("failed to decode request body: %w", err)
		writeError(w, err, http.StatusBadRequest)
		return
	}

	protected := checkProtected(req.Keys)
	if len(protected) > 0 {
		err = fmt.Errorf("the following environment variables are protected, and cannot be modified: % v", protected)
		writeError(w, err, http.StatusUnprocessableEntity)
		return
	}

	deleted := make([]string, 0, len(req.Keys))
	for _, k := range req.Keys {
		if _, ok := s.environ.Get(k); ok {
			deleted = append(deleted, k)
			delete(s.environ, k)
		}
	}

	json.NewEncoder(w).Encode(EnvDeleteResponse{Deleted: deleted})
	w.WriteHeader(http.StatusOK)
}

func checkProtected(candidates []string) []string {
	protected := make([]string, 0, len(candidates))
	for _, c := range candidates {
		if _, ok := agent.ProtectedEnv[c]; ok {
			protected = append(protected, c)
		}
	}
	return protected
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}
