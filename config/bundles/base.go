package bundles

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/spf13/viper"
)

// FindBundlesInConfig returns the list of bundles from the config file
func FindBundlesInConfig() []string {
	if err := config.InitViper(); err != nil {
		return []string{}
	}

	return viper.GetStringSlice("bundles")
}

// IsInstalled takes a bundle name and check if this name is in the config installed bundles
func IsInstalled(name string) bool {
	bundles := FindBundlesInConfig()

	return helpers.ContainsString(bundles, name)
}
