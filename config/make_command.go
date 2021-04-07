package config

import (
	"os"

	"github.com/edwinvautier/go-cli/prompt/entity"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitMakeCmdConfig creates the needed config for the create command by prompting user and doing other actions
func InitMakeCmdConfig(config *MakeCmdConfig) error {
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
	config.Box = packr.New("makeEntityBox", "../templates/makeEntity")
	return entity.PromptUserForEntityFields(&config.Entity)
}

// AddModelToConfig set the new bundle to true in config after install
func AddModelToConfig(newEntity entity.NewEntity) error {
	workdir, err := os.Getwd()
	if err != nil {
		return err
	}
	
	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")
	
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	models := viper.Get("models").([]entity.NewEntity)
	models = append(models, newEntity)
	log.Info("Using config file : ", viper.ConfigFileUsed())
	viper.WriteConfig()

	return nil
}

// InstallCmdConfig is the struct used to configure make command
type MakeCmdConfig struct {
	GoPackageFullPath string
	Box 			*packr.Box
	Entity		entity.NewEntity
}
