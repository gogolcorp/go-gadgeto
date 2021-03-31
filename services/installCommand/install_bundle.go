package installCommand

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InstallBundle(name string) error {
	box := packr.New("Bundles", "../../bundles")

	if(!bundleExists(name, box)) {
		return errors.New(name + " bundle does not exist")
	}

	// load & update config
	if err := updateConfig(name); err != nil {
		return err
	}

	// execute templates
	if err := executeTemplates(box, name); err != nil {
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
	if err := exec.Command("sh", "install.sh").Run(); err != nil {
		return err
	}

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

func executeTemplates(box *packr.Box, name string) error {
	
	return nil
}