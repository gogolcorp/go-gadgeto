package createCommand

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/gobuffalo/packd"
	//log "github.com/sirupsen/logrus"
)

func generateTemplates(config config.CreateCmdConfig) error {
	
	config.Box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name())
		GenerateFile(fileParts.Path, fileParts.Name, fileParts.OutputName, config)
		return nil
	})

	return nil
}