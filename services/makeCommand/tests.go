package makeCommand

import (
	"github.com/edwinvautier/go-gadgeto/config"
	log "github.com/sirupsen/logrus"
)

// MakeTests get config, and execute templates for the make fixtures command
func MakeTests(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if err := config.InitMakeTestsCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	if err := executeTestsTemplates(makeCmdConfig); err != nil {
		return err
	}

	log.Info("Tests created for ", modelName)
	log.Info("You can run these tests by running the make test command, don't forget to add tests cases")

	return nil
}
