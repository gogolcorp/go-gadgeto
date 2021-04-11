package installCommand

import (
	"github.com/edwinvautier/go-cli/config"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/files"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
)

func executeTemplates(name string, installCmdConfig config.InstallCmdConfig) error {
	installCmdConfig.Box = packr.New("templates", "../../bundles/"+name+"/templates")

	installCmdConfig.Box.Walk(func(path string, f packd.File) error {
		fInfo, _ := f.FileInfo()
		fileParts := helpers.GetFilePartsFromName(fInfo.Name(), "")
		files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, installCmdConfig)
		return nil
	})

	return nil
}
