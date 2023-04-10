package server

import (
	"net/http"

	"github.com/nint8835/dieppe/pkg/config"
)

type Server struct {
	router *http.ServeMux
	config *config.Config
}

func (s *Server) Serve() error {
	bindAddr := ":8080"
	if s.config.Server != nil && s.config.Server.BindAddr != nil {
		bindAddr = *s.config.Server.BindAddr
	}

	return http.ListenAndServe(bindAddr, s.router)
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	return &Server{
		router: mux,
		config: cfg,
	}
}
