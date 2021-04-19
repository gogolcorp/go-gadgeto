package makeCommand

import (
	"io/ioutil"
	"os"
	"strings"
	log "github.com/sirupsen/logrus"
	"github.com/edwinvautier/go-cli/services/filesystem"
)

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
		if (strings.Contains(line, "Db.AutoMigrate")) {
			line = line[:len(line) - 1] + ", &models." + modelNamePascalCase + "{})"
		}
		log.Info(line)
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
