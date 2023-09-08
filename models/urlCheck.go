package models

import "net/url"

// URLCheck if URL is valid
func URLCheck(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	// check if has scheme（suck as http, https）
	if u.Scheme == "" {
		return false
	}

	// check if it has host
	if u.Host == "" {
		return false
	}

	return true
}
