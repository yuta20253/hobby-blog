package storage

import "mime/multipart"

type FileStorage interface {
	Save(file *multipart.FileHeader) (path string, filename string, err error)
	Delete(path string) error
}
