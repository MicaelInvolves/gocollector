package utils

import "strings"

func IsValidValue(value string) bool {
	trim := strings.Trim(value, " ")
	return len(trim) > 0
}
