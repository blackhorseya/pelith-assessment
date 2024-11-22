package cmd

import (
	"github.com/blackhorseya/pelith-assessment/cmd/server"
	"github.com/blackhorseya/pelith-assessment/internal/shared/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
}

func init() {
	startCmd.PersistentFlags().String("token", "", "The token for Etherscan API")
	_ = viper.BindPFlag("services.server.etherscan.api_key", startCmd.PersistentFlags().Lookup("token"))

	startCmd.AddCommand(cmdx.NewServiceCmd("server", "Start the server", server.NewCmd))

	rootCmd.AddCommand(startCmd)
}
