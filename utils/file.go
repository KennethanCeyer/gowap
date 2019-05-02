package utils

import (
	"os"
)

func Exists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err != nil {
		return true
	}
	return !os.IsNotExist(err)
}

func CreateIfEmpty(dirPath string) (error, bool) {
	if Exists(dirPath) {
		return nil, false
	}

	return os.Mkdir(dirPath, 0755), true
}
