package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// AskGitUsername asks the user for a string used as the app name
func AskGitUsername(userName *string) error {
	prompt := &survey.Input{
		Message: "Please enter your git username",
	}
	return survey.AskOne(prompt, userName)
}