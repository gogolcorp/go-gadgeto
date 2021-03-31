package installCommand

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InstallBundle(name string) error {	
	var installCmdConfig config.InstallCmdConfig
	config.InitInstallCmdConfig(&installCmdConfig)
	box := packr.New("Bundles", "../../bundles")

	if(!bundleExists(name, box)) {
		return errors.New(name + " bundle does not exist")
	}

	// load & update config
	if err := updateConfig(name); err != nil {
		return err
	}

	// execute templates
	if err := executeTemplates(name, installCmdConfig); err != nil {
		return err
	}
	
	// if it exists, execute the shell script it contains
	if err := executeInstallScript(box, name); err != nil {
		return err
	}

	return nil
}

func executeInstallScript(box *packr.Box, name string) error {
	workdir, err := os.Getwd()

	if err != nil {
		return err
	}

	fileString, err := box.FindString("/" + name + "/install.sh")
	if err != nil {
		return err
	}

	ioutil.WriteFile(workdir + "/install.sh", []byte(fileString), 0755)
	exec.Command("sh", "install.sh").Run()

	return filesystem.RemoveSingle(workdir + "/install.sh")
}

func bundleExists(name string, box *packr.Box) bool {
	files := box.List()
	
	for _, file := range files {
		bundleName := strings.Split(file, "/")[0]
		if bundleName == name {
			return true
		}
	}

	return false
}

func updateConfig(name string) error {
	workdir, err := os.Getwd()

	if err != nil {
		return err
	}

	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")
	viper.Set("bundles." + name, true)
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file : ", viper.ConfigFileUsed())
		viper.WriteConfig()
		return nil
	}
	
	return err
}

func executeTemplates(name string, installCmdConfig config.InstallCmdConfig) error {
	box := packr.New("templates", "../../bundles/" + name + "/templates")
	
	box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name())
		GenerateFile(fileParts.Path, fileParts.Name, fileParts.OutputName, installCmdConfig, box)
		return nil
	})

	return nil
}

// GenerateFile takes params such as path, file name and desired name + config to generate the file after being templated
func GenerateFile(path string, name string, outputName string, installCmdConfig config.InstallCmdConfig, box *packr.Box) {
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
