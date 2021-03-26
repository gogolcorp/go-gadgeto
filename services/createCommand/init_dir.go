package createCommand

import (
	"os"

	"github.com/edwinvautier/go-cli/config"
	log "github.com/sirupsen/logrus"
)

// InitProject creates the directory for a new project and all needed structure depending on the config given
func InitProject(config *config.CreateCmdConfig) error {
	workingDirectory := getWorkDir()
	log.Info(workingDirectory)

	return nil
}

func getWorkDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't get your working directory")
	}

	return path
}
