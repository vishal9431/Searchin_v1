package utils;

import (
	"net/url"
	"strings"
)


func Contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func IsUrl(link string) bool {
	parsedUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}
	return strings.HasPrefix(parsedUrl.Scheme, "http")
}