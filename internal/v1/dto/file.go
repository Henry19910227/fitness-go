package dto

import "mime/multipart"

type File struct {
	FileNamed string
	Data multipart.File
}
