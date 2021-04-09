package config

import (
	"os"

	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/prompt/entity"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitMakeCmdConfig creates the needed config for the create command by prompting user and doing other actions
func InitMakeCmdConfig(config *MakeCmdConfig) error {
	configBase := initBasicConfig(config)

	config.GoPackageFullPath = configBase.PackagePath
	config.ProjectPath = configBase.ProjectPath
	config.Box = packr.New("makeEntityBox", "../templates/makeEntity")
	config.Entity.NameLowerCase = helpers.LowerCase(config.Entity.Name)
	config.Entity.NamePascalCase = helpers.UpperCaseFirstChar(config.Entity.Name)
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
	viper.SetDefault("models", map[string]map[string]string{})
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	models := viper.GetStringMap("models")
	models[newEntity.Name] = newEntity
	viper.Set("models", models)
	log.Info("Using config file : ", viper.ConfigFileUsed())
	viper.WriteConfig()

	return nil
}

// MakeCmdConfig is the struct used to configure make command
type MakeCmdConfig struct {
	GoPackageFullPath string
	Box               *packr.Box
	Entity            entity.NewEntity
	ProjectPath       string
}

// GetBox returns the box in which templates for make command are stored
func (cmd MakeCmdConfig) GetBox() *packr.Box {
	return cmd.Box
}

// GetProjectPath returns the path to project in user's computer
func (cmd MakeCmdConfig) GetProjectPath() string {
	return cmd.ProjectPath
}
