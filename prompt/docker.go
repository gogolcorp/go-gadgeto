package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

func AskToUseDocker(wantsDocker *bool) error {
	prompt := &survey.Confirm{
    Message: "Do you want to use docker ?",
	}
	return survey.AskOne(prompt, wantsDocker)
}