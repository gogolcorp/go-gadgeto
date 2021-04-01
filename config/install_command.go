package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
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

// UpdateConfigAfterInstalling set the new bundle to true in config after install
func UpdateConfigAfterInstalling(name string) error {
	workdir, err := os.Getwd()

	if err != nil {
		return err
	}

	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")
	viper.Set("bundles."+name, true)
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file : ", viper.ConfigFileUsed())
		viper.WriteConfig()
		return nil
	}

	return err
}

// InstallCmdConfig is the struct for the templates config of install command
type InstallCmdConfig struct {
	GoPackageFullPath string
}
