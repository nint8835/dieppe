package config

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// ParseFile parses the config contained within a given file.
func ParseFile(path string) (*Config, error) {
	parser := hclparse.NewParser()

	fileHcl, diags := parser.ParseHCLFile(path)
	//TODO: detect width, disable colour for non-tty
	wr := hcl.NewDiagnosticTextWriter(os.Stderr, parser.Files(), 80, true)
	if diags.HasErrors() {
		wr.WriteDiagnostics(diags)
		return nil, diags
	}

	var configFile Config
	if diags := gohcl.DecodeBody(fileHcl.Body, nil, &configFile); diags.HasErrors() {
		wr.WriteDiagnostics(diags)
		return nil, diags
	}

	return &configFile, nil
}
