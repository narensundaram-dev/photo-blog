package utils

import (
	"os"
	fp "path/filepath"
	"strings"
)

// MediaPath where the photos will be stored
var MediaPath = "media"

// GetRelativeMediaPath return the relative media path
func GetRelativeMediaPath(paths ...string) string {
	return fp.Join(MediaPath, strings.Join(paths, "/"))
}

// GetMediaPath returns the media path
func GetMediaPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return fp.Join(path, MediaPath)
}

// MakeSureThePath creates the media dir if doesn't exist
func MakeSureThePath(path string) {
	os.MkdirAll(path, os.ModePerm)
}
