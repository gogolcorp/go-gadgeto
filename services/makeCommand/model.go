package makeCommand

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/filesystem"
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

// AddModelToMigrations adds the model to migrations file
func AddModelToMigrations(modelNamePascalCase string) error {
	workdir := filesystem.GetWorkdirOrDie()
	migrationsFile, err := ioutil.ReadFile(workdir + "/shared/database/migrations.go")
	if err != nil {
		return err
	}
	migrationsContent := string(migrationsFile)

	var finalMigrationsLines []string
	migrationsLines := strings.Split(migrationsContent, "\n")

	for _, line := range migrationsLines {
		if strings.Contains(line, "Db.AutoMigrate") {
			line = line[:len(line)-1] + ", &models." + modelNamePascalCase + "{})"
		}
		finalMigrationsLines = append(finalMigrationsLines, line)
	}

	file, err := os.OpenFile(workdir+"/shared/database/migrations.go", os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(finalMigrationsLines, "\n"))

	return err
}
