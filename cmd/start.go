package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"github.com/nint8835/dieppe/pkg/server"
)

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the proxy server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(cfg)

		err := srv.Serve()
		if err != nil {
			slog.Error("failed to start server", "err", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCommand)
}
