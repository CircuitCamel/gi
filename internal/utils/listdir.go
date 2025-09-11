package utils

import (
	"os"
)

func ListDir(dirname string) ([]byte, error) {
	dirContent, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	var files string

	for _, v := range dirContent {
		files += v.Name()[:len(v.Name())-10] + "\n"
	}
	return []byte(files), nil
}
