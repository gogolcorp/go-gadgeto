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
	projectPath := workingDirectory + "/" + config.AppName

	if err := createProjectDir(projectPath); err != nil {
		return err
	}
	log.WithField("path", projectPath).Info("project directory created")
	
	if err := services.GitInit(projectPath); err != nil {
		return err
	}
	log.Info("git initialized")

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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	log.Warn("A directory with this name already exists.")

	wantsOverride := false
	prompt.AskToOverride(&wantsOverride)

	if wantsOverride {
		return removeAll(path)
	}

	return errors.New("Couldn't create project directory")
}

func removeAll(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.Mkdir(path, os.ModePerm)
}