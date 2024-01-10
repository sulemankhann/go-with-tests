package iteration

import (
	"strings"
)

// Repeat takes character and repat it by the count passed
func Repeat(character string, repeatCount int) string {
	repeated := strings.Repeat(character, repeatCount)

	return repeated
}
