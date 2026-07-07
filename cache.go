package main

import (
	"errors"
	"os"
	"path"
	"strings"
)


func _fsItemExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !(err != nil && errors.Is(err, os.ErrNotExist))
}

func BasePath() string{
  homedir, _ :=os.UserHomeDir()
  return path.Join(homedir,".cache/pokefetch")
}

func readyCache() {
	if !_fsItemExists(BasePath()) {
		os.Mkdir(BasePath(), os.FileMode(0522))
	}
}

func getFilePathForRequestUrl(url string) string {
	replacedUrl := strings.ReplaceAll(url, "/", "_")
	return path.Join(BasePath(), replacedUrl)
}

func validatePath(path string) error {
	if _fsItemExists(path) {
		_, err := os.Open(path)
		if err != nil {
			return errors.New("Error opening file")
		}
    return nil
	}
  return os.ErrNotExist
}

func readFileToBytes(validatedFilePath string) []byte {
	data, err := os.ReadFile(validatedFilePath)
	if err == nil {
		return data
	}
	return make([]byte, 0)
}

func writeRequestBytesToFile(requestBytes []byte, url string) error {
	cachePath := getFilePathForRequestUrl(url)
  err:=os.WriteFile(cachePath, requestBytes, 0644)
  return err
}

func getCachedFileContentOrError(url string) ([]byte, error) {
	readyCache()

	cachePath := getFilePathForRequestUrl(url)
  err := validatePath(cachePath)
	if err != nil {
		return make([]byte, 0), err
	}
  return readFileToBytes(cachePath), nil
}
