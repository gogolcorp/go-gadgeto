package createCommand

import (
	"errors"
	"os"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/prompt"
	"github.com/edwinvautier/go-cli/services"
	log "github.com/sirupsen/logrus"
)

// InitProject creates the directory for a new project and all needed structure depending on the config given
func InitProject(config *config.CreateCmdConfig) error {
	workingDirectory := getWorkDir()
	config.ProjectPath = workingDirectory + "/" + config.AppName

	if err := createProjectDir(config.ProjectPath); err != nil {
		return err
	}
	log.WithField("path", config.ProjectPath).Info("project directory created")

	if err := services.GitInit(config.ProjectPath); err != nil {
		return err
	}
	log.Info("git initialized")

	if err := generateTemplates(*config); err != nil {
		return err
	}
	log.Info("project initialization finished!")

	return nil
}

func getWorkDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't get your working directory")
	}

	return path
}

func createProjectDir(path string) error {
	if !services.DirectoryExists(path) {
		return os.Mkdir(path, os.ModePerm)
	}
	log.Warn("A directory with this name already exists.")

	wantsOverride := false
	prompt.AskToOverride(&wantsOverride)

	if wantsOverride {
		return services.RemoveDirAndFiles(path)
	}

	return errors.New("Couldn't create project directory")
}
