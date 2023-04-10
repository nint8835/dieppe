package config

// Server contains the configuration for the Dieppe server itself.
type Server struct {
	BindAddr *string `hcl:"bind_addr"`
}

// Config is the top-level configuration structure.
type Config struct {
	Server *Server `hcl:"server,block"`
}
