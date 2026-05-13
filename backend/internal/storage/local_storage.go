package storage

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type LocalStogare struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStogare {
	return &LocalStogare{
		basePath: basePath,
	}
}

func (s *LocalStogare) Save(file *multipart.FileHeader) (string, string, error) {
	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", "", err
	}

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := "uploads/" + fileName

	src, err := file.Open()

	if err != nil {
		return "", "", err
	}

	defer src.Close()

	dst, err := os.Create(path)

	if err != nil {
		return "", "", err
	}

	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return "", "", err
	}

	return path, fileName, nil
}

func (s *LocalStogare) Delete(path string) error {
	return os.Remove(path)
}
