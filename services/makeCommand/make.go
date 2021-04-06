package makeCommand

import "github.com/edwinvautier/go-cli/config"

func MakeEntity(entityName string) error {
	var makeCmdConfig config.MakeCmdConfig
	if err := config.InitMakeCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeTemplates(makeCmdConfig); err != nil {
		return err
	}
	
	return nil
}