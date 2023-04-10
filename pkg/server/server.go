package server

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/config"
)

type Server struct {
	router *http.ServeMux
	config *config.Config
}

func (s *Server) Serve() error {
	slog.Info("Starting Dieppe", "bind_addr", *s.config.Server.BindAddr)

	return http.ListenAndServe(*s.config.Server.BindAddr, s.router)
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	return &Server{
		router: mux,
		config: cfg,
	}
}
