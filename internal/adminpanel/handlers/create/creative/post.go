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
	"time"
)

func CreativePost(db dbifaces.DataBaseProvider, imageUploader imguploaderifaces.ImageUploader) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		partnerID := context.FormValue("partner_internal_creative_id")
		offerIDStr := context.FormValue("offer_id")
		typeStr := context.FormValue("type")
		if partnerID == "" || offerIDStr == "" {
			return context.SendString(400, "Обязательные поля не заполнены")
		}

		offerID, err := strconv.ParseUint(offerIDStr, 10, 64)
		if err != nil {
			return context.SendString(400, "Некорректный ID оффера")
		}

		resource, err := context.FormFile("image")

		var resourceURL string
		var width, height int

		if err == nil {
			resourceFile, err := resource.Open()
			if err != nil {
				return context.SendString(400, "Ошибка загрузки содержимого файла изображения")
			}
			defer resourceFile.Close()

			buf, err := io.ReadAll(resourceFile)
			if err != nil {
				return context.SendString(400, "Ошибка при чтении файла изображения")
			}
			resourceURL, err = imageUploader.UploadImage(bytes.NewReader(buf), resource.Filename)
			if err != nil {
				return context.SendString(400, "Ошибка при получении URL изображения")
			}
			width, height, err = imageutils.GetImageDimensions(bytes.NewReader(buf))
			if err != nil {
				return context.SendString(400, "Ошибка при получении размеров изображения")
			}
		} else {
			resourceURL, width, height = "", 0, 0
		}

		newCreative := dbmodels.Creative{
			PartnerInternalCreativeID: partnerID,
			OfferID:                   uint(offerID),
			Type:                      typeStr,
			ResourceURL:               resourceURL,
			Width:                     width,
			Height:                    height,
			AddedAt:                   time.Now(),
			UpdatedAt:                 time.Now(),
		}

		if err := db.Create(contextBackground, &newCreative); err != nil {
			return context.SendString(500, "Ошибка при создании креатива")
		}

		var creatives []dbmodels.Creative
		if err := db.Find(contextBackground, &creatives); err != nil {
			return context.SendString(500, "Ошибка при загрузке креативов")
		}

		response := `
			<div id="creative-result" hx-swap-oob="true" style="color:green;">
				Креатив успешно добавлен!
			</div>

		`

		return context.SendString(200, response)
	}
}
