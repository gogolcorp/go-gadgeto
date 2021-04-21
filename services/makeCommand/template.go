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
		fileParts := helpers.GetFilePartsFromName(fInfo.Name(), makeCmdConfig.Model.NameLowerCase+".go")
		files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)
		return nil
	})

	return nil
}

func executeModelTemplate(makeCmdConfig config.MakeCmdConfig) error {
	fileParts := helpers.GetFilePartsFromName("api/models/models_template.go.template", makeCmdConfig.Model.NameLowerCase+".go")
	files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)

	return nil
}

func executeFixturesTemplates(makeCmdConfig config.MakeCmdConfig) error {
	// Generate base
	fileParts := helpers.GetFilePartsFromName("fixtures/fixtures.go.template", "fixtures.go")
	files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)

	// Generate fixture templates
	fileParts = helpers.GetFilePartsFromName("fixtures/single_fixture.go.template", makeCmdConfig.Model.NameLowerCase+".go")
	files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)

	fileParts = helpers.GetFilePartsFromName("fixtures/fixture.json.template", makeCmdConfig.Model.NameLowerCase+".json")
	files.Generate(fileParts.Path, fileParts.Name, fileParts.OutputName, makeCmdConfig)
	return nil
}
