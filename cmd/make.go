package cmd

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/edwinvautier/go-cli/services/makeCommand"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// makeCmd represents the install command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "make is used to create new files, for example for models",
	Long:  `make is used to create new files, for example for models, it creates your model file after prompting you for fields`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isAMakeCommand(args[0]) {
			log.Fatal(args[0], " is not a make command!")
		}
		makeCommand.MakeEntity(args[1])
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

func isAMakeCommand(commandName string) bool {
	return commandName == "entity"
}
