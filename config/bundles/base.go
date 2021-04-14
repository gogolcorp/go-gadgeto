package bundles

import (
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/filesystem"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func FindBundlesInConfig() []string {
	workdir := filesystem.GetWorkdirOrDie()
	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-cli-config")
	viper.ReadInConfig()
	log.Info(viper.GetStringSlice("bundles"))

	return viper.GetStringSlice("bundles")
}

func IsInstalled(name string) bool {
	bundles := FindBundlesInConfig()

	return helpers.ContainsString(bundles, name)
}