package logs

import (
	"os"
	"strings"
)

func MkDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// path/to/whatever does not exist
		// use MkdirAll for nested directory
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func MkDirByFilePath(filePath string) error {
	pathSplit := strings.Split(filePath, "/")
	if (len(pathSplit) > 2 && pathSplit[0] == ".") || (len(pathSplit) > 1 && pathSplit[0] != ".") {
		dirName := strings.Join(pathSplit[:len(pathSplit)-1], "/")
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			// path/to/whatever does not exist
			// use MkdirAll for nested directory
			err := os.MkdirAll(dirName, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
