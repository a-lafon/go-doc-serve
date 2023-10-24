package fileHandler

import (
	"os"
)

type FileReader struct{}

func (fr *FileReader) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
