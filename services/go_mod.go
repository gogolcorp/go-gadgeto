package services

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func GoModInit(path string) error {
	cmd := exec.Command("cd", path, "&&", "go", "mod", "init")
	_, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
