package utils

import (
	"fmt"
	"os"
	"strings"
)

func NewID(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}

func HasValue(input string) bool {
	return input != ""
}

func SanatizeName(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}

func BoolPtr(b bool) *bool {
	return &b
}

func EnsureDirectoryExists(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	return nil
}
