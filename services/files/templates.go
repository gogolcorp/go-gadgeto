package files

import (
	"os"
	"text/template"

	"github.com/edwinvautier/go-gadgeto/config"
	log "github.com/sirupsen/logrus"
)

// Generate takes params such as path, file name and desired name + config to generate the file after being templated
func Generate(path string, name string, outputName string, commandConfig config.CommandConfigInterface) {
	// Get template content as string
	templateString, err := commandConfig.GetBox().FindString(path + name)
	if err != nil {
		log.WithField("path", path+name).Error("couldn't get template")
		return
	}

	// Create the directory if not exist
	if _, err := os.Stat(commandConfig.GetProjectPath() + "/" + path); os.IsNotExist(err) {
		os.MkdirAll(commandConfig.GetProjectPath()+"/"+path, os.ModePerm)
	}

	err = executeTemplate(commandConfig, outputName, commandConfig.GetProjectPath()+"/"+path, templateString)
	if err != nil {
		log.WithField("path", path+outputName).Error("couldn't create file")
		return
	}
}

func executeTemplate(config config.CommandConfigInterface, outputName string, path string, templateString string) error {
	// Create the file
	file, err := os.Create(path + outputName)
	if err != nil {
		return err
	}

	// Execute template and write file
	parsedTemplate := template.Must(template.New("template").Parse(templateString))

	return parsedTemplate.Execute(file, config)
}
