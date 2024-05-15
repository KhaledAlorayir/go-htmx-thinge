package common

import "strings"

func FormatValidationErrors(errors error) []string {
	return strings.Split(errors.Error(), "; ")
}
