package config

import (
	"os"

	"github.com/spf13/viper"
)

// InitInstallCmdConfig creates the needed config for the create command by prompting user and doing other actions
func InitInstallCmdConfig(config *InstallCmdConfig) error {
	workdir, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	config.GoPackageFullPath = viper.GetString("package")
	
	return nil
}

type InstallCmdConfig struct {
	GoPackageFullPath string
}