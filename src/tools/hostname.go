package tools

import "regexp"

func ValidateHostName(str string) bool {
	var validator = regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)

	matches := validator.MatchString(str)

	validator = nil
	return matches
}
