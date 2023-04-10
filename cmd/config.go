package cmd

import (
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Output Dieppe configuration",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		pp.Print(cfg)
	},
}

func init() {
	rootCmd.AddCommand(configCommand)
}
