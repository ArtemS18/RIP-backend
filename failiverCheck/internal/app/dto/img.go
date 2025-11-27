package dto

import "io"

type ComponentImgCreateDTO struct {
	File        io.Reader
	FilePath    string
	FileSize    int64
	ContentType string
}
