package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// AskDBMS asks the user for the database management system he wants
func AskDBMS(dbms *string) error {
	prompt := &survey.Select{
		Message: "Please enter application's name",
		Options: []string{"postgres", "mysql"},
	}
	return survey.AskOne(prompt, dbms)
}
