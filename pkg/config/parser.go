package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// ParseFiles parses the config contained within a collection of files.
func ParseFiles(paths []string) (*Config, error) {
	parser := hclparse.NewParser()

	var hclFiles []*hcl.File

	for _, filePath := range paths {
		hclFile, diags := parser.ParseHCLFile(filePath)
		//TODO: detect width, disable colour for non-tty
		wr := hcl.NewDiagnosticTextWriter(os.Stderr, parser.Files(), 80, true)
		if diags.HasErrors() {
			wr.WriteDiagnostics(diags)
			return nil, diags
		}

		hclFiles = append(hclFiles, hclFile)
	}

	mergedFile := hcl.MergeFiles(hclFiles)

	//TODO: detect width, disable colour for non-tty
	wr := hcl.NewDiagnosticTextWriter(os.Stderr, parser.Files(), 80, true)

	var configFile Config
	if diags := gohcl.DecodeBody(mergedFile, nil, &configFile); diags.HasErrors() {
		wr.WriteDiagnostics(diags)
		return nil, diags
	}

	return &configFile, nil
}

// ParseDir parses all config files in a given directory.
func ParseDir(dirPath string) (*Config, error) {
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

	return ParseFiles(filePaths)
}

// Parse parses the config at a given path.
// If the path is a directory, all HCL in the directory will be parsed.
func Parse(cfgPath string) (*Config, error) {
	fileInfo, err := os.Stat(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("error stating config path: %w", err)
	}

	if fileInfo.IsDir() {
		return ParseDir(cfgPath)
	}

	return ParseFiles([]string{cfgPath})
}
