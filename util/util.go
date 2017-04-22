package util

import (
	"os"
	"path"
)

// GetCWD returns the path to the current working directory
func GetCWD() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return path.Dir(ex)
}
