package createCommand

import (
	"errors"
	"os"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/prompt"
	"github.com/edwinvautier/go-cli/services"
	"github.com/edwinvautier/go-cli/services/filesystem"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	createProjectConfig(workingDirectory+"/"+config.AppName, config)

	if err := generateTemplates(*config); err != nil {
		return err
	}
	log.Info("project initialization finished!")

	return CleanAllFiles(config)
}

func getWorkDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't get your working directory")
	}

	return path
}

func createProjectDir(path string) error {
	if !filesystem.DirectoryExists(path) {
		return os.Mkdir(path, os.ModePerm)
	}
	log.Warn("A directory with this name already exists.")

	wantsOverride := false
	prompt.AskToOverride(&wantsOverride)

	if wantsOverride {
		return filesystem.RemoveDirAndFiles(path)
	}

	return errors.New("Couldn't create project directory")
}

func createProjectConfig(workdir string, config *config.CreateCmdConfig) {
	_, err := os.Create(workdir + "/.go-cli-config.yaml")
	if err != nil {
		log.Error("Couldn't create project config : ", err)
	}

	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")

	// Set config defaults
	viper.Set("package", config.GoPackageFullPath)
	viper.Set("database", config.DBMS)
	viper.Set("use_docker", config.UseDocker)
	viper.SetDefault("bundles.authenticator", false)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file : ", viper.ConfigFileUsed())
	}
	viper.WriteConfig()
}
