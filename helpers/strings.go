package helpers

import (
	"strings"
	"unicode"
)

// JoinString takes a pointer to a string and modify this string in order to remove spaces and replace them by dashes
func JoinString(str string) string {
	slices := strings.Split(str, " ")
	if slices[len(slices)-1] == "" {
		slices = slices[:len(slices)-1]
	}

	return strings.Join(slices, "-")
}

// GetFilePartsFromName tries to get path and name from the file name, also try to find desired extension
func GetFilePartsFromName(name string, outputName string) FileParts {
	var fileParts FileParts

	slices := strings.Split(name, "/")
	fileParts.Path = strings.Join(slices[:len(slices)-1], "/") + "/"
	fileParts.Name = slices[len(slices)-1]
	slices = strings.Split(fileParts.Name, ".")

	if outputName == "" {
		if strings.Contains(fileParts.Name, ".") {
			fileParts.OutputName = strings.Join(slices[:len(slices)-1], ".")
		} else {
			fileParts.OutputName = fileParts.Name
		}
	} else {
		fileParts.OutputName = outputName
	}

	return fileParts
}

// FileParts contains the needed informations to execute template for a file
type FileParts struct {
	Name       string
	Path       string
	OutputName string
}

// UpperCaseFirstChar returns the input string with first letter capitalized
func UpperCaseFirstChar(word string) string {
	a := []rune(word)
	if len(a) < 1 {
		return ""
	}

	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

// LowerCase returns input string lowercased
func LowerCase(name string) string {
	return strings.ToLower(name)
}

// Check if a substring exists in strings slice
func ContainsString(slice []string, substr string) bool {
	for _, element := range slice {
		if substr == element {
			return true
		}
	}

	return false
}
