package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/config"
)

var logLevel string

var cfgPath string
var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "dieppe",
	Short: "Vanity URL proxy for Go packages and Docker images",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "debug", "level to use for logging")
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "dieppe.hcl", "path to load config from")
}

func initLogging() {
	logLevelMap := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	level, levelValid := logLevelMap[logLevel]

	if !levelValid {
		slog.Error("invalid log level", "level", logLevel)
		os.Exit(1)
	}

	logger := slog.HandlerOptions{Level: level}.NewTextHandler(os.Stderr)

	slog.SetDefault(slog.New(logger))
}

func initConfig() {
	var err error
	cfg, err = config.Parse(cfgPath)
	if err != nil {
		if !errors.Is(err, config.HCLParseError) {
			slog.Error("failed to parse config", "err", err)
		}

		os.Exit(1)
	}
}
