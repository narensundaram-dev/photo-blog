package utils

import (
	"os"
	fp "path/filepath"
)

// GetMediaPath to return to the media path
func GetMediaPath() (string, error) {
	if path, err := os.Getwd(); err != nil {
		return "", err
	} else {
		return fp.Join(path, "media"), nil
	}
}
