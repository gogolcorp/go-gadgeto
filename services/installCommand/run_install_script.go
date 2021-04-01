package installCommand

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/edwinvautier/go-cli/services/filesystem"
	"github.com/gobuffalo/packr/v2"
)

func executeInstallScript(box *packr.Box, name string) error {
	workdir, err := os.Getwd()

	if err != nil {
		return err
	}

	fileString, err := box.FindString("/" + name + "/install.sh")
	if err != nil {
		return err
	}

	ioutil.WriteFile(workdir+"/install.sh", []byte(fileString), 0755)
	exec.Command("sh", "install.sh").Run()

	return filesystem.RemoveSingle(workdir + "/install.sh")
}
