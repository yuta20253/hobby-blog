package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
	}
}

func (s *LocalStorage) Save(reader io.Reader, filename string) (string, error) {
	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", err
	}

	safeName := filepath.Base(filename)

	filePath := filepath.Join(s.basePath, safeName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, reader); err != nil {
		_ = os.Remove(filePath)
		return "", err
	}

	return "/uploads/" + safeName, nil
}

func (s *LocalStorage) Delete(path string) error {
	fullPath := filepath.Join(s.basePath, filepath.Base(path))
	return os.Remove(fullPath)
}
