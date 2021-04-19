package makeCommand

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
)

// MakeModel creates  the config and execute templates in order to create a new Model
func MakeModel(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if config.IsInConfig(helpers.LowerCase(modelName)) {
		return updateModel(modelName)
	}

	if err := config.InitMakeModelCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeTemplates(makeCmdConfig); err != nil {
		return err
	}

	if err := AddModelToMigrations(makeCmdConfig.Model.NamePascalCase); err != nil {
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

	if err := executeTemplates(makeCmdConfig); err != nil {
		return err
	}

	return AddControllersToRouter(makeCmdConfig.Model.NamePascalCase)
}

func updateModel(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName

	if err := config.InitUpdateModelConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeModelTemplate(makeCmdConfig); err != nil {
		return err
	}

	return config.AddModelToConfig(makeCmdConfig.Model)
}
