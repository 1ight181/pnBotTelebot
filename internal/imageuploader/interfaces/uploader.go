package interfaces

import "io"

type ImageUploader interface {
	UploadImage(file io.Reader, filename string) (string, error)
}
