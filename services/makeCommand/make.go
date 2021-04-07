package makeCommand

import "github.com/edwinvautier/go-cli/config"

func MakeEntity(entityName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Entity.Name = entityName
	if err := config.InitMakeCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeTemplates(makeCmdConfig); err != nil {
		return err
	}

	return config.AddModelToConfig(makeCmdConfig.Entity)
}