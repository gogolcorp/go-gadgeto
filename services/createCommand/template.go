package createCommand

import (
	"os"
	"text/template"

	"github.com/edwinvautier/go-cli/config"
	log "github.com/sirupsen/logrus"
)

// GenerateFile takes params such as path, file name and desired name + config to generate the file after being templated
func GenerateFile(path string, name string, outputName string, config config.CreateCmdConfig) {
	// Get template content as string
	templateString, err := config.Box.FindString(path + name)
	if err != nil {
		log.Error(err)
		return
	}

	// Create the directory if not exist
	if _, err := os.Stat(config.ProjectPath + "/" + path); os.IsNotExist(err) {
		os.MkdirAll(config.ProjectPath+"/"+path, os.ModePerm)
	}

	err = executeTemplate(config, outputName, config.ProjectPath+"/"+path, templateString)
	if err != nil {
		log.Error(err)
		return
	}
}

func executeTemplate(config config.CreateCmdConfig, outputName string, path string, templateString string) error {
	// Create the file
	file, err := os.Create(path + outputName)
	if err != nil {
		log.Error(err)
		return err
	}

	// Execute template and write file
	parsedTemplate := template.Must(template.New("template").Parse(templateString))
	err = parsedTemplate.Execute(file, config)

	return nil
}
