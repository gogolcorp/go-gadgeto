package bundles

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/spf13/viper"
)

func FindBundlesInConfig() []string {
	if err := config.InitViper(); err != nil {
		return []string{}
	}

	return viper.GetStringSlice("bundles")
}

func IsInstalled(name string) bool {
	bundles := FindBundlesInConfig()

	return helpers.ContainsString(bundles, name)
}
