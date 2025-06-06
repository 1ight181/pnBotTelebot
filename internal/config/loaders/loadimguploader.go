package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadImageUploaderConfig(imageUploaderConfig conf.ImageUploader) string {
	freeimagehostApi := imageUploaderConfig.FreeimagehostApi

	return freeimagehostApi
}
