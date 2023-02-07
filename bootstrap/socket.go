package bootstrap

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
)

// SocketClient connects to the Job API.
type SocketClient struct {
	client *http.Client
	token  string
}

// NewSocketClient creates a new SocketClient.
func NewSocketClient() (*SocketClient, error) {
	sock := os.Getenv("BUILDKITE_AGENT_BOOTSTRAP_SOCK")
	if sock == "" {
		return nil, errors.New("BUILDKITE_AGENT_BOOTSTRAP_SOCK empty or undefined")
	}
	token := os.Getenv("BUILDKITE_AGENT_BOOTSTRAP_SOCK_TOKEN")
	if token == "" {
		return nil, errors.New("BUILDKITE_AGENT_BOOTSTRAP_SOCK_TOKEN empty or undefined")
	}
	return &SocketClient{
		client: &http.Client{
			Transport: &http.Transport{
				// Ignore arguments, dial socket
				DialContext: func(context.Context, string, string) (net.Conn, error) {
					return net.Dial("unix", sock)
				},
			},
		},
		token: token,
	}, nil
}

// Do attaches the token credential to the request, and then sends it to the
// socket.
func (c *SocketClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.token)
	return c.client.Do(req)
}
