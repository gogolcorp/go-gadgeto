package config

import (
	"os"

	"github.com/spf13/viper"
)

func initViper() error {
	workdir, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")

	return viper.ReadInConfig()
}
