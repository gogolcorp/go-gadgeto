package config

import (
	"os"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// CommandConfigInterface is the base for all go-cli commands
type CommandConfigInterface interface {
	GetBox() *packr.Box
	GetProjectPath() string
}

func initBasicConfig() baseConfig {
	if err := initViper(); err != nil {
		log.Error(err)
		return baseConfig{}
	}

	workdir, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return baseConfig{}
	}

	return baseConfig{
		PackagePath: viper.GetString("package"),
		ProjectPath: workdir,
	}
}

type baseConfig struct {
	PackagePath string
	ProjectPath string
}
