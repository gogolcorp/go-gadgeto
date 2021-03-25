package cmd

import (
	"github.com/edwinvautier/go-cli/prompt"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services"
	"os"
	"os/signal"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is used to initialize a new application.",
	Long:  `This command is used to initialize a new application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := CreateCmdConfig{
			Args: args,
		}
		initCmdConfig(&config)

		/*
		path, err := os.Getwd()
		if err != nil {
			log.Fatal("Couldn't find the current directory.")
		}

		services.CreateStructure(path+"/"+appName, modules, username, appName)
		*/
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Exiting...")
		os.Exit(1)
	}()
}

func initCmdConfig(config *CreateCmdConfig) {
	config.AppName = getAppName(config.Args)
	config.GitUserName = getGitUsername()
	config.DBMS = getDBMS()
	config.UseDocker = chooseToUseDocker()
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

type CreateCmdConfig struct {
	AppName 		string
	GitUserName string
	DBMS				string
	UseDocker		bool

	Args				[]string
}