package utils

import "strings"

func NewID(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}

func HasValue(input string) bool {
	return input != ""
}

func SanatizeName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}
