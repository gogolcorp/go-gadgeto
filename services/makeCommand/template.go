package makeCommand

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/files"
	"github.com/gobuffalo/packd"
)

func executeTemplates(makeCmdConfig config.MakeCmdConfig) error {

	makeCmdConfig.Box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name(), makeCmdConfig.Entity.NameLowerCase+".go")
		files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)
		return nil
	})

	return nil
}
