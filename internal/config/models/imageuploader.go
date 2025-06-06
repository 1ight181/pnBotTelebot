package models

import (
	"errors"
)

type ImageUploader struct {
	FreeimagehostApi string `mapstructure:"freeimagehost_api"`
}

func (iu *ImageUploader) Validate() error {
	if iu.FreeimagehostApi == "" {
		return errors.New("требуется указание api freeimage.host")
	}

	return nil
}
