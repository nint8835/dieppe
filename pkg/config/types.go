package config

// Server contains the configuration for the Dieppe server itself.
type Server struct {
	Host     string  `hcl:"host"`
	BindAddr *string `hcl:"bind_addr"`
}

type GoModule struct {
	ID       string  `hcl:",label"`
	Path     string  `hcl:"path"`
	Upstream string  `hcl:"upstream"`
	VCSType  *string `hcl:"vcs_type"`
}

func (m *GoModule) ImportPath(s *Server) string {
	return s.Host + "/" + m.Path
}

// Config is the top-level configuration structure.
type Config struct {
	Server *Server `hcl:"server,block"`

	GoModules []*GoModule `hcl:"go_module,block"`
}
