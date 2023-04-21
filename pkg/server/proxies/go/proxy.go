package goproxy

import (
	"bytes"
	_ "embed"
	"html/template"
	"net/http"

	"github.com/yuin/goldmark"
	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/config"
)

type proxyRespCtx struct {
	ImportPath  string
	VCSType     string
	UpstreamURL string
	Readme      template.HTML
	Module      *config.GoModule
}

//go:embed proxy_resp.gohtml
var respTmplBody string

var respTmpl = template.Must(template.New("proxy_resp").Parse(respTmplBody))

type GoProxy struct {
	config *config.Config
	module *config.GoModule
}

func (p *GoProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var readmeBody bytes.Buffer

	if err := goldmark.Convert([]byte(p.module.Readme), &readmeBody); err != nil {
		slog.Error("Failed to render module readme", "module", p.module.ID, "error", err)
		readmeBody.Reset()
	}

	_ = respTmpl.Execute(w, proxyRespCtx{
		ImportPath:  p.module.ImportPath(p.config.Server),
		VCSType:     *p.module.VCSType,
		UpstreamURL: p.module.Upstream,
		Readme:      template.HTML(readmeBody.String()),
		Module:      p.module,
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
