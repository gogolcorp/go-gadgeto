package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// AskApplicationName asks the user for a string used as the app name
func AskApplicationName(appName *string) error {
	prompt := &survey.Input{
		Message: "Please enter application's name",
	}
	return survey.AskOne(prompt, appName)
}