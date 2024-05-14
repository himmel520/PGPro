package app

import (
	"context"
	"net/http"
	"time"

	"github.com/himmel520/pgPro/internal/config"
)

// Server defines an HTTP server.
type Server struct {
	httpServer *http.Server
}

// New creates a new instance of Server with the specified cfg and handler.
func New(cfg config.ServerConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

// Run starts the HTTP server.
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting active connections.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
