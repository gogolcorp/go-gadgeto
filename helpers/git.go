package helpers

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetGitUsername tries to find the git username inside cli config or in the git config
func GetGitUsername() string {
	// try to get from viper
	userName := viper.GetString("git-username")
	if userName != "" {
		return userName
	}

	// else try to get from the command line
	cmd := exec.Command("git", "config", "user.name")
	stdout, err := cmd.Output()
	userName = string(stdout)

	if err != nil {
		log.Error(err)
		return ""
	}

	// Store to config if not empty
	if userName != "" {
		viper.Set("git-username", userName)
		viper.WriteConfig()
	}

	return userName
}