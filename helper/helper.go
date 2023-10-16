package helper

import "net/url"

func IsValidURL(testURL string) bool {
	u, err := url.Parse(testURL)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}
