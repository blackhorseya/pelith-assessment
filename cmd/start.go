package cmd

import (
	"github.com/blackhorseya/pelith-assessment/cmd/server"
	"github.com/blackhorseya/pelith-assessment/pkg/cmdx"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.AddCommand(cmdx.NewServiceCmd("server", "Start the server", server.NewCmd))
}
