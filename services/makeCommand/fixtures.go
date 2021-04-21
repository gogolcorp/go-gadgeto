package makeCommand

import (
	"github.com/edwinvautier/go-cli/config"
	log "github.com/sirupsen/logrus"
)

// MakeFixtures get config, and execute templates for the make fixtures command
func MakeFixtures(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if err := config.InitMakeFixturesCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	// Templates : base(package + env & makefile ?) & each model fixtures
	if err := executeFixturesTemplates(makeCmdConfig); err != nil {
		return err
	}

	log.Info("Fixtures created for ", modelName)

	return nil
}
