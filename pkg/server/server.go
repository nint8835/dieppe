package server

import (
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/config"
	"github.com/nint8835/dieppe/pkg/server/proxies/go"
)

func withLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug(
			"Request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
		)
		handler.ServeHTTP(w, r)
	})
}

type Server struct {
	router *http.ServeMux
	config *config.Config
}

func (s *Server) Serve() error {
	slog.Info("Starting Dieppe", "bind_addr", *s.config.Server.BindAddr)

	return http.ListenAndServe(*s.config.Server.BindAddr, withLogging(s.router))
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	goproxy.Register(cfg, mux)

	return &Server{
		router: mux,
		config: cfg,
	}
}
