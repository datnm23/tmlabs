package pkg

import (
	"os"
)

func CheckFileExisted(path string) (bool, error) {
	_, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}

func CheckFolderExisted(path string) (bool, error) {
	_, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
