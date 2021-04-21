package config

import (
	"github.com/edwinvautier/go-gadgeto/services/filesystem"
	"github.com/spf13/viper"
)

// InitViper init acces to viper config in workdir
func InitViper() error {
	workdir := filesystem.GetWorkdirOrDie()
	viper.AddConfigPath(workdir)
	viper.SetConfigName(".go-gadgeto-config")

	return viper.ReadInConfig()
}
