package jobapi

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/buildkite/agent/v3/env"
)

// Server is a Job API server. It provides an HTTP API with which to interact with the job currently running in the buildkite agent
// and allows jobs to introspect and mutate their own state
type Server struct {
	SocketPath string
	environ    env.Environment
	token      string
	httpSvr    *http.Server
	started    bool
}

// NewServer creates a new Job API server
// socketPath is the path to the socket on which the server will listen
// environ is the environment which the server will mutate and inspect as part of its operation
func NewServer(socketPath string, environ env.Environment) (server *Server, token string, err error) {
	exists, err := socketExists(socketPath)
	if err != nil {
		return nil, "", err
	}

	if exists {
		err = os.RemoveAll(socketPath)
		if err != nil {
			return nil, "", fmt.Errorf("removing existing socket: %w", err)
		}
	}

	token, err = generateToken(32)
	if err != nil {
		return nil, "", fmt.Errorf("generating token: %w", err)
	}

	return &Server{
		SocketPath: socketPath,
		environ:    environ,
		token:      token,
	}, token, nil
}

// Start starts the server in a goroutine, returning an error if the server can't be started
func (s *Server) Start() error {
	if s.started {
		return errors.New("server already started")
	}

	r := s.router()
	l, err := net.Listen("unix", s.SocketPath)
	if err != nil {
		return fmt.Errorf("listening on socket: %w", err)
	}

	s.httpSvr = &http.Server{Handler: r}
	go func() {
		_ = s.httpSvr.Serve(l)
	}()
	s.started = true

	return nil
}

// Stop gracefully shuts the server down, blocking until all requests have been served or the grace period has expired
// It returns an error if the server has not been started
func (s *Server) Stop() error {
	if !s.started {
		return errors.New("server not started")
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	defer serverStopCtx()

	// REVIEW: Should we capture signals here?

	// Shutdown signal with grace period of 30 seconds
	shutdownCtx, _ := context.WithTimeout(serverCtx, 10*time.Second)

	go func() {
		<-shutdownCtx.Done()
		if shutdownCtx.Err() == context.DeadlineExceeded {
			// REVIEW: What should we do in this situation? Force a return? Log something?
		}
	}()

	// Trigger graceful shutdown
	err := s.httpSvr.Shutdown(shutdownCtx)
	if err != nil {
		return fmt.Errorf("shutting down Job API server: %w", err)
	}

	err = os.Remove(s.SocketPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing socket: %w", err)
	}

	return nil
}

func socketExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("checking if socket exists: %w", err)
}

func generateToken(len int) (string, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("reading from random: %w", err)
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
