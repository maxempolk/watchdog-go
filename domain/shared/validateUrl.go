package shared

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidateURL(urlStr string) (string, error) {
	urlStr = strings.TrimSpace(urlStr)
	if urlStr == "" {
		return "", fmt.Errorf("URL is empty")
	}

	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "https://" + urlStr
	}

	parsed, err := url.Parse(urlStr)
	if err != nil || parsed.Host == "" {
		return "", fmt.Errorf("invalid URL: %s", urlStr)
	}

	return urlStr, nil
}