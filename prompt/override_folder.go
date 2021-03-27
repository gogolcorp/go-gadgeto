package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// AskToOverride simply asks the user if the app should be dockerized or not
func AskToOverride(wantsOverride *bool) error {
	prompt := &survey.Confirm{
		Message: "Override folder ?",
	}
	return survey.AskOne(prompt, wantsOverride)
}
