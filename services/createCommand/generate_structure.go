package createCommand

import (
	"github.com/edwinvautier/go-gadgeto/config"
	"github.com/edwinvautier/go-gadgeto/helpers"
	"github.com/edwinvautier/go-gadgeto/services/files"
	"github.com/gobuffalo/packd"
	//log "github.com/sirupsen/logrus"
)

func generateTemplates(config config.CreateCmdConfig) error {
	config.Box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name(), "")
		files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, config)
		return nil
	})

	return nil
}
