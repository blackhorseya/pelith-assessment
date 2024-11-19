package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
}

func init() {
	rootCmd.AddCommand(startCmd)
}
