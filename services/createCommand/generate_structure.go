package createCommand

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/templates"
)

func generateTemplates(config config.CreateCmdConfig) error {
	templates.GenerateFile("", "main.txt", "main.go", config)
	return nil
}