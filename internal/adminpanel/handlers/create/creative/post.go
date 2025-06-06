package creative

import (
	"bytes"
	ctx "context"
	"io"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	imguploaderifaces "pnBot/internal/imageuploader/interfaces"
	imageutils "pnBot/internal/imageutils"
	"strconv"
)

func CreativePost(db dbifaces.DataBaseProvider, imageUploader imguploaderifaces.ImageUploader) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		partnerID := context.FormValue("partner_internal_creative_id")
		offerIDStr := context.FormValue("offer_id")
		typeStr := context.FormValue("type")
		if partnerID == "" || offerIDStr == "" {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Обязательные поля не заполнены</div>")
		}

		offerID, err := strconv.ParseUint(offerIDStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Некорректный ID оффера</div>")
		}

		resource, err := context.FormFile("image")

		var resourceURL string
		var width, height int

		if err == nil {
			resourceFile, err := resource.Open()
			if err != nil {
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка загрузки содержимого файла изображения</div>")
			}
			defer resourceFile.Close()

			buf, err := io.ReadAll(resourceFile)
			if err != nil {
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при чтении файла изображения</div>")
			}
			resourceURL, err = imageUploader.UploadImage(bytes.NewReader(buf), resource.Filename)
			if err != nil {
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при получении URL изображения</div>")
			}
			width, height, err = imageutils.GetImageDimensions(bytes.NewReader(buf))
			if err != nil {
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при получении размеров изображения</div>")
			}
		} else {
			resourceURL, width, height = "", 0, 0
		}

		newCreative := dbmodels.Creative{
			PartnerInternalCreativeId: partnerID,
			OfferId:                   uint(offerID),
			Type:                      typeStr,
			ResourceUrl:               resourceURL,
			Width:                     width,
			Height:                    height,
		}

		if err := db.Create(contextBackground, &newCreative); err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при создании креатива</div>")
		}

		var creatives []dbmodels.Creative
		if err := db.Find(contextBackground, &creatives); err != nil {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при загрузке креативов</div>")
		}

		response := `
			<div class="success-box"> Креатив успешно добавлен! </div>
		`

		return context.Status(200).Type("text/html").SendString(response)
	}
}
