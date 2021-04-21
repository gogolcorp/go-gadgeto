package makeCommand

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/services/filesystem"
	log "github.com/sirupsen/logrus"
)

// MakeFixtures get config, and execute templates for the make fixtures command
func MakeFixtures(modelName string) error {
	var makeCmdConfig config.MakeCmdConfig
	makeCmdConfig.Model.Name = modelName
	if err := config.InitMakeFixturesCmdConfig(&makeCmdConfig); err != nil {
		return err
	}

	workdir := filesystem.GetWorkdirOrDie()
	if _, err := os.Stat(workdir + "/fixtures/fixtures.go"); os.IsNotExist(err) {
		if err := AddFixturesToMain(makeCmdConfig); err != nil {
			log.Info("Don't forget to call fixtures when you need", err)
		}
		AddFixturesToEnv()
	}

	// Templates : base(package + env & makefile ?) & each model fixtures
	if err := executeFixturesTemplates(makeCmdConfig); err != nil {
		return err
	}

	log.Info("Fixtures created for ", modelName)
	log.Info("You can execute fixtures at project startup by setting env variable RUN_FIXTURES to true")

	return nil
}

// AddFixturesToMain add lines to call fixtures package inside main file of the project
func AddFixturesToMain(config config.MakeCmdConfig) error {
	workdir := filesystem.GetWorkdirOrDie()
	mainFile, err := ioutil.ReadFile(workdir + "/main.go")
	if err != nil {
		return err
	}
	mainContent := string(mainFile)

	var finalMainLines []string
	mainLines := strings.Split(mainContent, "\n")
	// Create string
	linesToAppend := []string{
		"\t\"" + config.GoPackageFullPath + "/fixtures\"",
		"\tif env.GoDotEnvVariable() == \"true\" {\n\t\tfixtures.RunFixtures()\n\t}",
	}
	for _, line := range mainLines {
		finalMainLines = append(finalMainLines, line)
		if strings.Contains(line, "Migrate()") {
			finalMainLines = append(finalMainLines, linesToAppend[1])
		} else if strings.Contains(line, "import (") {
			finalMainLines = append(finalMainLines, linesToAppend[0])
		}
	}

	file, err := os.OpenFile(workdir+"/main.go", os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(finalMainLines, "\n"))

	return err
}

// AddFixturesToEnv writes the RUN_FIXTURES variable inside .env.dist file
func AddFixturesToEnv() error {
	workdir := filesystem.GetWorkdirOrDie()
	envFile, err := ioutil.ReadFile(workdir + "/.env.dist")
	if err != nil {
		return err
	}
	envContent := string(envFile)

	envLines := strings.Split(envContent, "\n")
	envLines = append(envLines, "\nRUN_FIXTURES=false")

	file, err := os.OpenFile(workdir+"/.env.dist", os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(envLines, "\n"))

	return err
}
