package modelPrompt

import (
	"io/ioutil"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/filesystem"
	log "github.com/sirupsen/logrus"
)

// ModelField represents a single field from a model
type ModelField struct {
	Type      string
	Name      string
	IsSlice   bool
	SliceType string
}

// NewModel represents the full model that user wants to create
type NewModel struct {
	Name           string
	NamePascalCase string
	NameLowerCase  string
	HasDate        bool
	HasCustomTypes bool
	Fields         []ModelField
}

// PromptUserForModelFields prompts user in the CLI to choose model fields wanted
func PromptUserForModelFields(model *NewModel) error {
	for {
		fieldName := ""
		if err := promptForFieldName(&fieldName); err != nil {
			return err
		}

		// If field name is empty then stop the function
		if fieldName == "" {
			break
		}

		var fieldType string
		if err := promptForFieldType(&fieldType); err != nil {
			return err
		}

		field := ModelField{
			Name:    helpers.UpperCaseFirstChar(fieldName),
			Type:    fieldType,
			IsSlice: false,
		}

		if field.Type == "date" {
			model.HasDate = true
		} else if field.Type == "slice" {
			field.IsSlice = true

			sliceTypePrompt := &survey.Select{
				Message: "Slice type :",
				Options: GetTypeOptions(),
			}
			survey.AskOne(sliceTypePrompt, &field.SliceType)

			if choosedCustomType(field.SliceType) {
				model.HasCustomTypes = true
			}
		}

		if choosedCustomType(field.Type) {
			model.HasCustomTypes = true
		}

		model.Fields = append(model.Fields, field)
	}

	return nil
}

// promptForFieldName asks the user for a string used as the app name
func promptForFieldName(fieldName *string) error {
	prompt := &survey.Input{
		Message: "Choose new field name (Press enter to stop adding fields)",
	}
	return survey.AskOne(prompt, fieldName)
}

func promptForFieldType(fieldType *string) error {
	typePrompt := &survey.Select{
		Message: "Choose type :",
		Options: GetTypeOptions(),
	}
	return survey.AskOne(typePrompt, fieldType)
}

// GetTypeOptions returns a list of strings for user prompt of data types when creating new models
func GetTypeOptions() []string {
	modelsList := GetModelsList()
	options := []string{"string", "boolean", "int", "uint", "float32", "float64", "date", "slice"}
	options = append(options, modelsList...)

	return options
}

// GetModelsList returns a slice of strings with all the models names found in the models/ dir
func GetModelsList() []string {
	workdir := filesystem.GetWorkdirOrDie()
	files, err := ioutil.ReadDir(workdir + "/api/models")
	if err != nil {
		log.Fatal(err)
	}
	models := make([]string, 0)
	for _, file := range files {
		name := helpers.UpperCaseFirstChar(strings.Split(file.Name(), ".go")[0])
		models = append(models, name)
	}

	return models
}

func choosedCustomType(cType string) bool {
	modelsList := GetModelsList()
	for _, modelName := range modelsList {
		if modelName == cType {
			return true
		}
	}

	return false
}
