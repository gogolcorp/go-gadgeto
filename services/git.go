package services

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetGitUsername tries to find the git username inside cli config or in the git config
func GetGitUsername() string {
	var userName string

	// try to get from viper
	userName = viper.GetString("git-username")
	if userName != "" {
		return userName
	}

	// Try to get from git config
	userName = getFromGit()
	storeToConfig(userName)

	return userName
}

func getFromGit() string {
	cmd := exec.Command("git", "config", "user.name")
	stdout, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(stdout)
}

func storeToConfig(userName string) {
	if userName != "" {
		viper.Set("git-username", userName)
		viper.WriteConfig()
	}
}