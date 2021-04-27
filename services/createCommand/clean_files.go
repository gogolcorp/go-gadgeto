package createCommand

import (
	"github.com/edwinvautier/go-gadgeto/config"
	"github.com/edwinvautier/go-gadgeto/services/filesystem"
	log "github.com/sirupsen/logrus"
)

// CleanAllFiles removed all files that are not necessary as the user chose in the config
func CleanAllFiles(config *config.CreateCmdConfig) error {
	if !config.UseDocker {
		if err := removeDockerFiles(config); err != nil {
			log.Error("Could'nt remove docker files: ", err)
		}
	}

	return nil
}

func removeDockerFiles(config *config.CreateCmdConfig) error {
	// Remove docker folder
	if err := filesystem.RemoveDirAndFiles(config.ProjectPath + "/docker"); err != nil {
		return err
	}

	if err := filesystem.RemoveSingle(config.ProjectPath + "/docker"); err != nil {
		return err
	}

	// Remove docker compose
	return filesystem.RemoveSingle(config.ProjectPath + "/docker-compose.yml")
}
