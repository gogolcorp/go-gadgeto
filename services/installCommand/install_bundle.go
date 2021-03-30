package installCommand

import (
	"errors"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InstallBundle(name string) error {
	if(!bundleExists(name)) {
		return errors.New(name + " bundle does not exist")
	}

	// load & update config
	if err := updateConfig(name); err != nil {
		return err
	}

	// if it exists, execute the shell script it contains

	// execute templates

	return nil
}

func bundleExists(name string) bool {
	box := packr.New("Bundles", "../../bundles")
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