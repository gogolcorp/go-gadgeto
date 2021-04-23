package createCommand

import (
	"errors"
	"os"

	"github.com/edwinvautier/go-gadgeto/config"
	"github.com/edwinvautier/go-gadgeto/prompt"
	"github.com/edwinvautier/go-gadgeto/services"
	"github.com/edwinvautier/go-gadgeto/services/filesystem"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitProject creates the directory for a new project and all needed structure depending on the config given
func InitProject(config *config.CreateCmdConfig) error {
	workdir := filesystem.GetWorkdirOrDie()
	config.ProjectPath = workdir + "/" + config.AppName

	if err := createProjectDir(config.ProjectPath); err != nil {
		return err
	}
	log.WithField("path", config.ProjectPath).Info("project directory created")

	if err := services.GitInit(config.ProjectPath); err != nil {
		return err
	}
	log.Info("git initialized")

	createProjectConfig(workdir+"/"+config.AppName, config)

	if err := generateTemplates(*config); err != nil {
		return err
	}
	log.Info("project initialization finished!")

	return CleanAllFiles(config)
}

func createProjectDir(path string) error {
	if !filesystem.DirectoryExists(path) {
		return os.Mkdir(path, os.ModePerm)
	}
	log.Warn("a directory with this name already exists.")

	wantsOverride := false
	prompt.AskToOverride(&wantsOverride)

	if wantsOverride {
		return filesystem.RemoveDirAndFiles(path)
	}

	return errors.New("couldn't create project directory")
}

func createProjectConfig(workdir string, config *config.CreateCmdConfig) {
	_, err := os.Create(workdir + "/.go-gadgeto-config.yaml")
	if err != nil {
		log.Error("couldn't create project config : ", err)
	}

	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-gadgeto-config")

	// Set config defaults
	viper.Set("package", config.GoPackageFullPath)
	viper.Set("database", config.DBMS)
	viper.Set("use_docker", config.UseDocker)
	viper.SetDefault("bundles", []string{})

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("using config file : ", viper.ConfigFileUsed())
	}
	viper.WriteConfig()
}
