package util

import "strings"

// IsURL takes a string path and returns
// true if it is a URL
func IsURL(path string) bool {
	return strings.Contains(path, "http")
}
