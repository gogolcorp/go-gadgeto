package createCommand

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/services/filesystem"
	log "github.com/sirupsen/logrus"
)

// CleanAllFiles removed all files that are not necessary as the user chose in the config
func CleanAllFiles(config *config.CreateCmdConfig) error {
	if (!config.UseDocker) {
		if err := removeDockerFiles(config); err != nil {
			log.Error("Could'nt remove docker files", err)
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

func removeAuthenticationFiles(config *config.CreateCmdConfig) error {
	// Remove controller
	if err := filesystem.RemoveSingle(config.ProjectPath + "/api/controllers/authentication.go"); err != nil {
		return err
	}

	// Remove middleware
	if err := filesystem.RemoveSingle(config.ProjectPath + "/shared/middlewares/authorization_middleware.go"); err != nil {
		return err
	}

	// Remove password hasher service
	if err := filesystem.RemoveSingle(config.ProjectPath + "/shared/services/token.go"); err != nil {
		return err
	}

	// Remove token service
	return filesystem.RemoveSingle(config.ProjectPath + "/shared/services/password_hasher.go")
}