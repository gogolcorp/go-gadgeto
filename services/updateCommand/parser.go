package updateCommand

import (
	"strings"

	"github.com/edwinvautier/go-cli/prompt/entity"
)

// ParseEntity takes a file string and try to parse entity from it
func ParseEntity(model *entity.NewEntity, fileContent string) {
	lines := strings.Split(fileContent, "\n")
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
			var field entity.EntityField
			parseField(model, &field, line)

			if len(field.Name) > 2 || len(field.Type) > 2 {
				model.Fields = append(model.Fields, field)
			}
		}

		if strings.Contains(line, "type") {
			lineIsStruct = true
		}
	}
}

func assignType(model *entity.NewEntity, field *entity.EntityField, element string) {
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

func haveFoundFields(model *entity.NewEntity) bool {
	return len(model.Fields) > 0
}

func parseField(model *entity.NewEntity, field *entity.EntityField, line string) {
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
