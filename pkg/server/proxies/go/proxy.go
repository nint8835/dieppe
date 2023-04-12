package goproxy

import (
	_ "embed"
	"html/template"
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/config"
)

type proxyRespCtx struct {
	ImportPath  string
	VCSType     string
	UpstreamURL string
}

//go:embed proxy_resp.gohtml
var respTmplBody string

var respTmpl = template.Must(template.New("proxy_resp").Parse(respTmplBody))

type GoProxy struct {
	config *config.Config
	module *config.GoModule
}

func (p *GoProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = respTmpl.Execute(w, proxyRespCtx{
		ImportPath:  p.module.ImportPath(p.config.Server),
		VCSType:     *p.module.VCSType,
		UpstreamURL: p.module.Upstream,
	})
}

func Register(cfg *config.Config, mux *http.ServeMux) {
	for _, module := range cfg.GoModules {
		proxy := &GoProxy{
			config: cfg,
			module: module,
		}

		proxyPath := "/" + module.Path
		mux.Handle(proxyPath, proxy)

		slog.Debug("Registered Go module proxy", "module", module.ID, "path", proxyPath)
	}
}
