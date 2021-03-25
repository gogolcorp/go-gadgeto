package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// AskToUseDocker simply asks the user if the app should be dockerized or not
func AskToUseDocker(wantsDocker *bool) error {
	prompt := &survey.Confirm{
    Message: "Do you want to use docker ?",
	}
	return survey.AskOne(prompt, wantsDocker)
}