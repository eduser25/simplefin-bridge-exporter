package config

import (
	"fmt"
	"net/url"
	"os"
)

func ReadAndDeleteAccessURLFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading AccessUrl path: %v", err)
	}

	parsedUrl, err := url.Parse(string(data))
	if err != nil {
		return "", fmt.Errorf("error parsing AccessUrl: %v", err)
	}

	err = os.Remove(filepath)
	if err != nil {
		return "", fmt.Errorf("error removing AccessUrl file: %v", err)
	}

	return parsedUrl.String(), nil
}
