package server

import (
	_ "embed"
	"html/template"
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/config"
	"github.com/nint8835/dieppe/pkg/server/proxies/go"
)

type indexCtx struct {
	Config *config.Config
}

//go:embed index.gohtml
var indexTmplBody string

var indexTmpl = template.Must(template.New("index").Parse(indexTmplBody))

func withBranding(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Dieppe")
		handler.ServeHTTP(w, r)
	})
}

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

	return http.ListenAndServe(*s.config.Server.BindAddr, withBranding(withLogging(s.router)))
}

func (s *Server) ServeIndex(w http.ResponseWriter, r *http.Request) {
	_ = indexTmpl.Execute(w, indexCtx{
		Config: s.config,
	})
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	goproxy.Register(cfg, mux)

	srv := &Server{
		router: mux,
		config: cfg,
	}

	mux.HandleFunc("/", srv.ServeIndex)

	return srv
}
