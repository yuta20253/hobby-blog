package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
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

	filePath := filepath.Join(s.basePath, s.generateFileName(filename))

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, reader); err != nil {
		_ = os.Remove(filePath)
		return "", err
	}

	return filePath, nil
}

func (s *LocalStorage) Delete(path string) error {
	return os.Remove(path)
}

func (s *LocalStorage) generateFileName(filename string) string {
	ext := filepath.Ext(filename)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}
