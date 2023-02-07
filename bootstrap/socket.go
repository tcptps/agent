package bootstrap

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
)

func ConnectToSocket() (*http.Client, string, error) {
	sock := os.Getenv("BUILDKITE_AGENT_BOOTSTRAP_SOCK")
	if sock == "" {
		return nil, "", errors.New("BUILDKITE_AGENT_BOOTSTRAP_SOCK empty or undefined")
	}
	token := os.Getenv("BUILDKITE_AGENT_BOOTSTRAP_SOCK_TOKEN")
	if token == "" {
		return nil, "", errors.New("BUILDKITE_AGENT_BOOTSTRAP_SOCK_TOKEN empty or undefined")
	}
	return &http.Client{
		Transport: &http.Transport{
			// Ignore arguments, dial socket
			DialContext: func(context.Context, string, string) (net.Conn, error) {
				return net.Dial("unix", sock)
			},
		},
	}, token, nil
}
