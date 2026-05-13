package storage

import "io"

type FileStorage interface {
	Save(reader io.Reader, filename string) (path string, err error)
	Delete(path string) error
}
