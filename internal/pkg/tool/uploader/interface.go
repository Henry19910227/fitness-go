package uploader

import "io"

type Tool interface {
	Save(file io.Reader, fileNamed string) (string, error)
	Delete(fileNamed string) error
}

type Size interface {
	Size() int64
}
