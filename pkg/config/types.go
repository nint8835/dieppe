package config

// Server contains the configuration for the Dieppe server itself.
type Server struct {
	Host     string  `hcl:"host"`
	BindAddr *string `hcl:"bind_addr"`
}

// Link defines a link to be displayed on a package's page.
type Link struct {
	Text string `hcl:"text"`
	URL  string `hcl:"url"`
}

// GoModule defines a Go module that should be proxied.
type GoModule struct {
	ID          string  `hcl:"id,label"`
	DisplayName *string `hcl:"display_name"`
	Description string  `hcl:"description,optional"`
	Path        string  `hcl:"path"`
	Upstream    string  `hcl:"upstream"`
	VCSType     *string `hcl:"vcs_type"`

	Readme string `hcl:"readme,optional"`
	Links  []Link `hcl:"link,block"`
}

func (m *GoModule) ImportPath(s Server) string {
	return s.Host + "/" + m.Path
}

// Config is the top-level configuration structure.
type Config struct {
	Server Server `hcl:"server,block"`

	GoModules []*GoModule `hcl:"go_module,block"`
}
