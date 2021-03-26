package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/services/createCommand"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is used to initialize a new application.",
	Long:  `This command is used to initialize a new application.`,
	Run: func(cmd *cobra.Command, args []string) {
		commandConfig := config.CreateCmdConfig{
			Args: args,
		}
		config.InitCreateCmdConfig(&commandConfig)
		
		createCommand.InitProject(&commandConfig)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Exiting...")
		os.Exit(1)
	}()
}