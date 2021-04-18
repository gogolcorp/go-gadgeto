package config

import (
	"errors"

	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/prompt/modelPrompt"
	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/gobuffalo/packr/v2"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitMakeModelCmdConfig creates the needed config for the create command by prompting user and doing other actions
func InitMakeModelCmdConfig(config *MakeCmdConfig) error {
	configBase := initBasicConfig()

	config.GoPackageFullPath = configBase.PackagePath
	config.ProjectPath = configBase.ProjectPath
	config.Box = packr.New("makeModelBox", "../templates/makeModel")
	config.Model.NameLowerCase = helpers.LowerCase(config.Model.Name)
	config.Model.NamePascalCase = helpers.UpperCaseFirstChar(config.Model.Name)
	return modelPrompt.PromptUserForModelFields(&config.Model)
}

// AddModelToConfig set the new bundle to true in config after install
func AddModelToConfig(newModel modelPrompt.NewModel) error {
	workdir := filesystem.GetWorkdirOrDie()

	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")
	viper.SetDefault("models", map[string]map[string]string{})
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	models := viper.GetStringMap("models")
	models[newModel.Name] = newModel
	viper.Set("models", models)
	log.Info("Using config file : ", viper.ConfigFileUsed())
	viper.WriteConfig()

	return nil
}

// MakeCmdConfig is the struct used to configure make command
type MakeCmdConfig struct {
	GoPackageFullPath string
	Box               *packr.Box
	Model             modelPrompt.NewModel
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

// InitMakeCRUDCmdConfig creates the needed config for the make crud command
func InitMakeCRUDCmdConfig(config *MakeCmdConfig) error {
	// Get model from config
	if err := InitViper(); err != nil {
		log.Fatal("couldn't read config, try again")
	}
	modelsStructs := viper.GetStringMap("models")
	var models []string
	for key := range modelsStructs {
		models = append(models, key)
	}
	if !helpers.ContainsString(models, config.Model.Name) {
		return errors.New("model does not exist, check if you api/models/" + config.Model.Name + ".go file exists and try running go-cli update")
	}

	modelData := viper.GetStringMap("models." + config.Model.Name)
	var model modelPrompt.NewModel
	if err := mapstructure.Decode(modelData, &model); err != nil {
		return errors.New("error while decoding " + config.Model.Name)
	}
	config.Model = model
	configBase := initBasicConfig()

	config.GoPackageFullPath = configBase.PackagePath
	config.ProjectPath = configBase.ProjectPath
	config.Box = packr.New("makeCRUDBox", "../templates/makeCRUD")

	return nil
}
