package imageutils

import (
	"image"
	_ "image/jpeg" // регистрируем JPEG-декодер
	_ "image/png"  // регистрируем PNG-декодер
	"io"
)

func GetImageDimensions(r io.Reader) (width, height int, err error) {
	img, _, err := image.DecodeConfig(r) // только конфиг, без полной декодировки
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
