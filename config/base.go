package config

import (
	"github.com/edwinvautier/go-gadgeto/services/filesystem"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
)

// CommandConfigInterface is the base for all go-gadgeto commands
type CommandConfigInterface interface {
	GetBox() *packr.Box
	GetProjectPath() string
}

func initBasicConfig() baseConfig {
	if err := InitViper(); err != nil {
		return baseConfig{}
	}
	workdir := filesystem.GetWorkdirOrDie()

	return baseConfig{
		PackagePath: viper.GetString("package"),
		ProjectPath: workdir,
	}
}

type baseConfig struct {
	PackagePath string
	ProjectPath string
}
