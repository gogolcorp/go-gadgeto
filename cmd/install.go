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
	"github.com/edwinvautier/go-cli/config/bundles"
	"github.com/edwinvautier/go-cli/services/installCommand"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install bundles to your app",
	Long:  `A command that install bundles from edwinvautier/go-cli/bundles`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, bundleName := range args {
			if bundles.IsInstalled(bundleName) {
				continue
			}
			if err := installCommand.InstallBundle(bundleName); err != nil {
				log.Error(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
