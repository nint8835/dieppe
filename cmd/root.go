package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:   "dieppe",
	Short: "Vanity URL proxy for Go packages and Docker images ",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogging)

	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "debug", "level to use for logging")
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
		log.Fatalf("invalid log level: %s", logLevel)
	}

	logger := slog.HandlerOptions{Level: level}.NewTextHandler(os.Stderr)

	slog.SetDefault(slog.New(logger))
}
