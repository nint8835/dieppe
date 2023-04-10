package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// HCLParseError occurs when an error is encountered while parsing HCL.
// It should not be logged to the console, as the error message will be printed by the HCL parser.
var HCLParseError = errors.New("failed to parse config")

func parseFiles(paths []string) (*Config, error) {
	parser := hclparse.NewParser()

	var hclFiles []*hcl.File

	for _, filePath := range paths {
		hclFile, diags := parser.ParseHCLFile(filePath)
		//TODO: detect width, disable colour for non-tty
		wr := hcl.NewDiagnosticTextWriter(os.Stderr, parser.Files(), 80, true)
		if diags.HasErrors() {
			wr.WriteDiagnostics(diags)
			return nil, HCLParseError
		}

		hclFiles = append(hclFiles, hclFile)
	}

	mergedFile := hcl.MergeFiles(hclFiles)

	//TODO: detect width, disable colour for non-tty
	wr := hcl.NewDiagnosticTextWriter(os.Stderr, parser.Files(), 80, true)

	var configFile Config
	if diags := gohcl.DecodeBody(mergedFile, nil, &configFile); diags.HasErrors() {
		wr.WriteDiagnostics(diags)
		return nil, HCLParseError
	}

	populateDefaults(&configFile)

	return &configFile, nil
}

func parseDir(dirPath string) (*Config, error) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error listing directory: %w", err)
	}

	var filePaths []string
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		if !strings.HasSuffix(dirEntry.Name(), ".hcl") {
			continue
		}

		filePaths = append(filePaths, path.Join(dirPath, dirEntry.Name()))
	}

	return parseFiles(filePaths)
}

func populateDefaults(config *Config) {
	if config.Server == nil {
		config.Server = &Server{}
	}

	if config.Server.BindAddr == nil {
		config.Server.BindAddr = new(string)
		*config.Server.BindAddr = ":8080"
	}
}

// Parse parses the config at a given path.
// If the path is a directory, all HCL in the directory will be parsed.
func Parse(cfgPath string) (*Config, error) {
	fileInfo, err := os.Stat(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("error stating config path: %w", err)
	}

	if fileInfo.IsDir() {
		return parseDir(cfgPath)
	}

	return parseFiles([]string{cfgPath})
}
