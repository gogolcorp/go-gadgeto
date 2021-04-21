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

// MakeCmdConfig is the struct used to configure make command
type MakeCmdConfig struct {
	GoPackageFullPath string
	Box               *packr.Box
	Model             modelPrompt.NewModel
	ProjectPath       string
	FixturesModels    []string
}

// GetBox returns the box in which templates for make command are stored
func (cmd MakeCmdConfig) GetBox() *packr.Box {
	return cmd.Box
}

// GetProjectPath returns the path to project in user's computer
func (cmd MakeCmdConfig) GetProjectPath() string {
	return cmd.ProjectPath
}

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

// InitMakeCRUDCmdConfig creates the needed config for the make crud command
func InitModelConfig(config *MakeCmdConfig) error {
	// Get model from config
	if err := InitViper(); err != nil {
		log.Fatal("couldn't read config, try again")
	}

	if !IsInConfig(config.Model.Name) {
		return errors.New("could'nt find this model, run go-cli update or go-cli make model to fix it")
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

	return nil
}

// InitMakeCRUDCmdConfig inits a config for the make CRUD command
func InitMakeCRUDCmdConfig(config *MakeCmdConfig) error {
	if err := InitModelConfig(config); err != nil {
		return err
	}
	config.Box = packr.New("makeCRUDBox", "../templates/makeCRUD")

	return nil
}

// InitMakeFixturesCmdConfig inits a config for the make fixtures command
func InitMakeFixturesCmdConfig(config *MakeCmdConfig) error {
	if err := InitModelConfig(config); err != nil {
		return err
	}
	config.Box = packr.New("makeFixturesBox", "../templates/makeFixtures")
	config.FixturesModels = filesystem.GetFixturesModelsList()

	if !helpers.ContainsString(config.FixturesModels, config.Model.NamePascalCase) {
		config.FixturesModels = append(config.FixturesModels, config.Model.NamePascalCase)
	}

	return nil
}

// IsInConfig returns a boolean telling wether the modelName was found in config or not
func IsInConfig(modelName string) bool {
	if err := InitViper(); err != nil {
		log.Error("error when loading viper")
		return false
	}
	modelsStructs := viper.GetStringMap("models")
	var models []string
	for key := range modelsStructs {
		models = append(models, key)
	}

	return helpers.ContainsString(models, modelName)
}

// InitUpdateModelConfig is the same as initMakeModelConfig but it gather infos from config first
func InitUpdateModelConfig(config *MakeCmdConfig) error {
	log.Info("model already exists! Select fields to add :")
	modelData := viper.GetStringMap("models." + config.Model.Name)
	var model modelPrompt.NewModel
	if err := mapstructure.Decode(modelData, &model); err != nil {
		return errors.New("error while decoding " + config.Model.Name)
	}
	config.Model = model

	return InitMakeModelCmdConfig(config)
}
