package config

import (
	"strings"

	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/prompt"
	"github.com/edwinvautier/go-cli/services"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitCreateCmdConfig creates the needed config for the create command by prompting user and doing other actions
func InitCreateCmdConfig(config *CreateCmdConfig) {
	config.AppName = getAppName(config.Args)
	config.GitUserName = getGitUsername()
	config.DBMS = getDBMS()
	config.UseDocker = chooseToUseDocker()
	config.GoPackageFullPath = "github.com/" + config.GitUserName + "/" + config.AppName
	config.Box = packr.New("My Box", "../templates")
}

func getAppName(args []string) string {
	appName := strings.Join(args, "-")
	// Check if the app name is empty
	if appName == "" {
		prompt.AskApplicationName(&appName)
		appName = helpers.JoinString(appName)
	}
	return appName
}

func getGitUsername() string {
	userName := services.GetGitUsername()
	if userName == "" {
		if err := prompt.AskGitUsername(&userName); err != nil {
			log.Fatal(err)
		}
		viper.Set("git-username", userName)
	}
	return userName
}

func getDBMS() string {
	// Get the desired DB management system
	dbms := ""
	if err := prompt.AskDBMS(&dbms); err != nil {
		log.Fatal(err)
	}
	return dbms
}

func chooseToUseDocker() bool {
	// Ask user wether to use docker or not
	wantsDocker := false
	prompt.AskToUseDocker(&wantsDocker)

	return wantsDocker
}

// CreateCmdConfig is the needed config for the command to work
type CreateCmdConfig struct {
	AppName           string
	GitUserName       string
	DBMS              string
	UseDocker         bool
	GoPackageFullPath string
	Args              []string
	Box               *packr.Box
	ProjectPath       string
}
