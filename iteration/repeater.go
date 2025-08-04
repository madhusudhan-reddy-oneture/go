package iteration

import "strings"

// Repeat returns string that has character repeated count times
func Repeat(character string, count uint) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(character)
	}
	return repeated.String()
}
