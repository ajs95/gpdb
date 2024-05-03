package fsutils

import (
	"errors"
	"os"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . FileSystem
type FileSystem interface {
	IsFilePresent(string) (bool, error)
}

func NewFileSystem() FileSystem {
	return &fileSystemUtil{}
}

type fileSystemUtil struct{}

func (fileSystemUtil) IsFilePresent(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
