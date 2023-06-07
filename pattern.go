package khanRouter

import (
	"log"
	"strings"
)

// getValidPattern returns a valid route pattern by cleaning and formatting the provided pattern string.
// It removes leading and trailing slashes, eliminates empty parts, and ensures a leading slash at the beginning.
// Example:
// getValidPattern("test/") returns "/test"
// getValidPattern("/test/") returns "/test"
// getValidPattern("test") returns "/test"
func getValidPattern(pattern string) string {
	if strings.TrimSpace(pattern) == "" {
		return ""
	}
	if strings.Contains(pattern, " ") {
		log.Panicf("httpRouter: invalid pattern '%s'", pattern)
	}
	pattern = strings.Trim(pattern, "/")
	parts := strings.Split(pattern, "/")
	cleanParts := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			cleanParts = append(cleanParts, part)
		}
	}
	return "/" + strings.Join(cleanParts, "/")
}
