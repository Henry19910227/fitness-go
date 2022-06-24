package file

import "mime/multipart"

type NamedField struct {
	Named string
}
type DataField struct {
	Data multipart.File
}
