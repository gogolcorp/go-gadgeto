package makeCommand

import "github.com/edwinvautier/go-cli/config"

// MakeModel creates  the config and execute templates in order to create a new Model
func MakeModel(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if err := config.InitMakeCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeTemplates(makeCmdConfig); err != nil {
		return err
	}

	return config.AddModelToConfig(makeCmdConfig.Model)
}
