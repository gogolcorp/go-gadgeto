package makeCommand

import (
	"os"

	"text/template"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/gobuffalo/packd"
	log "github.com/sirupsen/logrus"
)

func executeTemplates(makeCmdConfig config.MakeCmdConfig) error {

	makeCmdConfig.Box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name())
		generateFile(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)
		return nil
	})

	return nil
}

func generateFile(path string, name string, outputName string, makeCmdConfig config.MakeCmdConfig) {
	// Get template content as string
	templateString, err := makeCmdConfig.Box.FindString(path + name)
	if err != nil {
		log.Error(err)
		return
	}
	workdir, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}
	// Create the directory if not exist
	if _, err := os.Stat(workdir + "/" + path); os.IsNotExist(err) {
		os.MkdirAll(workdir+"/"+path, os.ModePerm)
	}

	err = executeTemplate(makeCmdConfig, outputName, workdir+"/"+path, templateString)
	if err != nil {
		log.Error(err)
		return
	}
}

func executeTemplate(makeCmdConfig config.MakeCmdConfig, outputName string, path string, templateString string) error {
	// Create the file
	file, err := os.Create(path + outputName)
	if err != nil {
		log.Error(err)
		return err
	}
	// Execute template and write file
	parsedTemplate := template.Must(template.New("template").Parse(templateString))
	err = parsedTemplate.Execute(file, makeCmdConfig)

	return nil
}
