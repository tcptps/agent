package jobapi

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type Server struct {
	SocketPath string
	token      string
	httpSvr    *http.Server
}

func NewServer(socketPath string) (server *Server, token string, err error) {
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
		token:      token,
	}, token, nil
}

func (s *Server) Start() error {
	r := s.router()
	l, err := net.Listen("unix", s.SocketPath)
	if err != nil {
		return fmt.Errorf("listening on socket: %w", err)
	}

	s.httpSvr = &http.Server{Handler: r}
	go func() {
		_ = s.httpSvr.Serve(l)
	}()

	return nil
}

func (s *Server) Stop() error {
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
