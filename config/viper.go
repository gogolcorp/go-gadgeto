package config

import (
	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/spf13/viper"
)

// init acces to viper config in workdir
func InitViper() error {
	workdir := filesystem.GetWorkdirOrDie()
	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")

	return viper.ReadInConfig()
}
