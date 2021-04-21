package filesystem

import (
	"io/ioutil"
	"strings"

	"github.com/edwinvautier/go-cli/helpers"
	log "github.com/sirupsen/logrus"
)

// GetModelsList returns a slice of strings with all the models names found in the models/ dir
func GetModelsList() []string {
	workdir := GetWorkdirOrDie()
	files, err := ioutil.ReadDir(workdir + "/api/models")
	if err != nil {
		log.Fatal(err)
	}

	models := make([]string, 0)
	for _, file := range files {
		name := helpers.UpperCaseFirstChar(strings.Split(file.Name(), ".go")[0])
		models = append(models, name)
	}

	return models
}

// GetFixturesModelsList returns a slice of strings with all the models names found in the fixtures/ dir
func GetFixturesModelsList() []string {
	workdir := GetWorkdirOrDie()
	files, err := ioutil.ReadDir(workdir + "/fixtures")
	models := make([]string, 0)
	if err != nil {
		log.Error("No fixtures", err)
		return models
	}

	for _, file := range files {
		name := helpers.UpperCaseFirstChar(strings.Split(file.Name(), ".go")[0])
		if !strings.Contains(name, "json") && !strings.Contains(name, "Fixture") {
			log.Info(name)
			models = append(models, name)
		}
	}

	return models
}
