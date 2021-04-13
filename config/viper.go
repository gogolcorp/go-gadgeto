package config

import (
	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/spf13/viper"
)

func initViper() error {
	workdir:= filesystem.GetWorkdirOrDie()
	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")

	return viper.ReadInConfig()
}
