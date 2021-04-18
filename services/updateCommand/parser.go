package updateCommand

import (
	"strings"

	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/prompt/modelPrompt"
)

// ParseModel takes a file string and try to parse model from it
func ParseModel(model *modelPrompt.NewModel, fileContent string) {
	content := strings.ReplaceAll(fileContent, "\t", " ")
	lines := strings.Split(content, "\n")

	// look inside it for infos
	lineIsStruct := false
	for _, line := range lines {
		if hasClosingBracket(line) {
			lineIsStruct = false
			if haveFoundFields(model) {
				break
			}
		}

		if lineIsStruct {
			var field modelPrompt.ModelField
			parseField(model, &field, line)

			if (len(field.Name) > 2 || len(field.Type) > 2) && helpers.LowerCase(field.Name) != "id" {
				model.Fields = append(model.Fields, field)
			}
		}

		if strings.Contains(line, "type") {
			lineIsStruct = true
		}
	}
}

func assignType(model *modelPrompt.NewModel, field *modelPrompt.ModelField, element string) {
	if strings.Contains(element, "[]") {
		field.IsSlice = true
		field.SliceType = strings.Trim(element, "[]")
		element = "slice"
	} else if strings.Contains(element, "Time") {
		element = "date"
		model.HasDate = true
	}
	field.Type = element
}

func hasClosingBracket(line string) bool {
	return strings.Contains(line, "}")
}

func haveFoundFields(model *modelPrompt.NewModel) bool {
	return len(model.Fields) > 0
}

func parseField(model *modelPrompt.NewModel, field *modelPrompt.ModelField, line string) {
	elements := strings.Split(line, " ")

	for _, element := range elements {
		element := strings.Trim(element, "\t")
		if len(element) < 2 {
			continue
		}
		if field.Name == "" {
			field.Name = element
		} else if field.Type == "" {
			assignType(model, field, element)
		}
	}
}
