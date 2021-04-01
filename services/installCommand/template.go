package installCommand

import (
	"os"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"text/template"
)

func executeTemplates(name string, installCmdConfig config.InstallCmdConfig) error {
	box := packr.New("templates", "../../bundles/"+name+"/templates")

	box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name())
		generateFile(fileParts.Path, fileParts.Name, fileParts.OutputName, installCmdConfig, box)
		return nil
	})

	return nil
}

func generateFile(path string, name string, outputName string, installCmdConfig config.InstallCmdConfig, box *packr.Box) {
	// Get template content as string
	templateString, err := box.FindString(path + name)
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

	err = executeTemplate(installCmdConfig, outputName, workdir+"/"+path, templateString)
	if err != nil {
		log.Error(err)
		return
	}
}

func executeTemplate(installCmdConfig config.InstallCmdConfig, outputName string, path string, templateString string) error {
	// Create the file
	file, err := os.Create(path + outputName)
	if err != nil {
		log.Error(err)
		return err
	}
	// Execute template and write file
	parsedTemplate := template.Must(template.New("template").Parse(templateString))
	err = parsedTemplate.Execute(file, installCmdConfig)

	return nil
}
