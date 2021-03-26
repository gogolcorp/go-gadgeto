package helpers

import (
	"strings"
)

// JoinString takes a pointer to a string and modify this string in order to remove spaces
func JoinString(str string) string {
	slices := strings.Split(str, " ")
	return strings.Join(slices, "-")
}

// GetFilePartsFromName tries to get path and name from the file name, also try to find desired extension
func GetFilePartsFromName(name string) FileParts {
	var fileParts FileParts

	slices := strings.Split(name, "/")
	fileParts.Name = slices[len(slices) - 1]
	fileParts.Path = strings.Join(slices[:len(slices)- 1], "/") + "/"
	
	slices = strings.Split(fileParts.Name, ".")
	fileParts.OutputName = strings.Join(slices[:len(slices) - 1], ".")

	return fileParts
}

// FileParts contains the needed informations to execute template for a file
type FileParts struct {
	Name 				string
	Path 				string
	OutputName 	string
}