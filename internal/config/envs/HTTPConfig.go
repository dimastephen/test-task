package envs

import (
	"net"
	"os"
	"test-task/internal/config"
)

const (
	httpPort = "HTTP_PORT"
	httpHost = "HTTP_HOST"
)

type httpConfig struct {
	host string
	port string
}

func (h *httpConfig) Address() string {
	return net.JoinHostPort(h.host, h.port)
}

func NewHTTPConfig() config.HTTPConfig {
	return &httpConfig{
		port: os.Getenv(httpPort),
		host: os.Getenv(httpHost),
	}
}
