package entity

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/edwinvautier/go-cli/helpers"
	log "github.com/sirupsen/logrus"
)


type EntityField struct {
	Type 			string
	Name 			string
	IsSlice 	bool
	SliceType	string
}

type NewEntity struct {
	Name 						string
	NamePascalCase 	string
	NameLowerCase 	string
	HasDate 				bool
	HasCustomTypes	bool
	Fields					[]EntityField
}

// PromptUserForEntityFields prompts user in the CLI to choose entity fields wanted
func PromptUserForEntityFields(entity *NewEntity) error{
	for true {
		fieldName := ""
		if err := promptForFieldName(&fieldName); err != nil {
			return err
		}

		// If field name is empty then stop the function
		if fieldName == "" {
			break
		}

		field := EntityField{
			Name: helpers.UpperCaseFirstChar(fieldName),
			IsSlice: false,
		}

		if err := promptForFieldType(&field.Type); err != nil {
			return err
		}

		if field.Type == "date" {
			entity.HasDate = true
		} else if field.Type == "slice" {
			field.IsSlice = true

			sliceTypePrompt := &survey.Select{
				Message: "Slice type :",
				Options: GetTypeOptions(),
			}
			survey.AskOne(sliceTypePrompt, &field.SliceType)

			if choosedCustomType(field.SliceType) {
				entity.HasCustomTypes = true
				field.SliceType = "models." + field.SliceType
			}
		}

		if choosedCustomType(field.Type) {
			entity.HasCustomTypes = true
			field.Type = "models." + field.Type
		}
		
		entity.Fields = append(entity.Fields, field)
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
		Message: "Choose a type for " + *fieldType + ":",
		Options: GetTypeOptions(),
	}
	return survey.AskOne(typePrompt, &fieldType)
}

// GetTypeOptions returns a list of strings for user prompt of data types when creating new models
func GetTypeOptions() []string {
	entitiesList := GetEntitiesList()
	options := []string{"string", "boolean", "int", "float", "date", "slice"}
	for _, entity := range entitiesList {
		options = append(options, entity)
	}

	return options
}

// GetEntitiesList returns a slice of strings with all the entities names found in the models/ dir
func GetEntitiesList() []string {
	workdir, err := os.Getwd()
	if err != nil {
		log.Error(err)
	}
	files, err := ioutil.ReadDir(workdir + "/api/models")
	if err != nil {
		log.Fatal(err)
	}
	entities := make([]string, 0)
	for _, file := range files {
		entities = append(entities, strings.Split(file.Name(), "Struct.go")[0])
	}

	return entities
}

func choosedCustomType(cType string) bool{
	entitiesList := GetEntitiesList()
	log.Info(cType)
	for _, entityName := range entitiesList {
		log.Info(entityName)
        if entityName == cType {
            return true
        }
	}
	
	return false
}