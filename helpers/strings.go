package helpers

import (
	"strings"
)

// JoinString takes a pointer to a string and modify this string in order to remove spaces
func JoinString(str string) string {
	slices := strings.Split(str, " ")
	return strings.Join(slices, "-")
}