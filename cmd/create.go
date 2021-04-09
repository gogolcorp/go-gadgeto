package cmd

/*
Copyright © 2021 Edwin Vautier <edwin.vautier@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

		if err := createCommand.InitProject(&commandConfig); err != nil {
			log.Error("create command failed : ", err)
		}
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
