package cmd

import (
	"os"

	"github.com/Ghazallghe/gofiber-monitoring/pkg/cmd/serve"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	//nolint: exhaustruct
	root := &cobra.Command{
		Use:   "monitor",
		Short: "final project of snapp course",
	}

	root.AddCommand(serve.New())

	if err := root.Execute(); err != nil {
		os.Exit(ExitFailure)
	}
}
