package utils

import "regexp"

func FilterAlphanumeric(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.ReplaceAllString(input, "")
}

func FilterUnsafeChars(input string) string {
	re := regexp.MustCompile(`[<>"'&;%$|*{}()\\/\[\]=]`)
	return re.ReplaceAllString(input, "")
}
