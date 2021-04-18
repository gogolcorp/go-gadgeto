package config

import (
	"io/ioutil"

	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/prompt/modelPrompt"
	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/edwinvautier/go-cli/services/updateCommand"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// UpdateConfig reads the project main parts in order to refresh the config store in the .go-cli-config.yml
func UpdateConfig() error {
	if err := InitViper(); err != nil {
		return err
	}

	// Update models
	return updateModels()
}

func updateModels() error {
	// Get config models list
	configModels := viper.GetStringMap("models")

	// Get project models list
	projectModels := modelPrompt.GetModelsList()

	// Trigger config generation for each new model
	for _, modelName := range projectModels {
		if configModels[helpers.LowerCase(modelName)] != nil {
			continue
		}

		var model modelPrompt.NewModel
		model.Name = modelName
		if err := generateModel(&model); err != nil {
			log.Error("couldn't add " + modelName + "to config")
			return err
		}

		if err := AddModelToConfig(model); err != nil {
			log.Error("couldn't add model to config : ", err)
			return err
		}
	}

	return nil
}

func generateModel(model *modelPrompt.NewModel) error {
	model.NamePascalCase = model.Name
	model.NameLowerCase = helpers.LowerCase(model.Name)

	// Get file content
	workdir := filesystem.GetWorkdirOrDie()
	filePath := workdir + "/api/models/" + model.NameLowerCase + ".go"
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Error("couldn't read : ", filePath)
		return err
	}
	contentString := string(content)
	updateCommand.ParseModel(model, contentString)

	return nil
}
