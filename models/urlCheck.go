package models

import (
	"net/url"
	"regexp"
)

// URLCheck checks if the given string contains a valid URL
func URLCheck(str string) bool {
	// Regular expression for identifying possible URL fragments
	urlRegexp := regexp.MustCompile(`https?://[^\s]+|[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all instances that might be URLs in the string
	matches := urlRegexp.FindAllString(str, -1)

	for _, match := range matches {
		if isValidURL(match) {
			return true
		} else {
			// Try prepending "https://" to incomplete URLs
			withHTTP := "https://" + match
			if isValidURL(withHTTP) {
				return true
			}
		}
	}
	return false
}

// isValidURL checks if a given string is a valid URL
func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
