package makeCommand

import (
	"github.com/edwinvautier/go-cli/config"
)

// MakeModel creates  the config and execute templates in order to create a new Model
func MakeModel(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if err := config.InitMakeModelCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeTemplates(makeCmdConfig); err != nil {
		return err
	}

	return config.AddModelToConfig(makeCmdConfig.Model)
}

// MakeCrud creates controllers for the model chosen
func MakeCrud(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if err := config.InitMakeCRUDCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	return executeTemplates(makeCmdConfig)
}
