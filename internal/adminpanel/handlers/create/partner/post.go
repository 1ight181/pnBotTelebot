package partner

import (
	"bytes"
	ctx "context"
	"io"
	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	imguploaderifaces "pnBot/internal/imageuploader/interfaces"
	"strconv"
	"time"
)

func PartnerPost(db dbifaces.DataBaseProvider, imageUploader imguploaderifaces.ImageUploader) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		var partners []dbmodels.Partner
		if err := db.Find(contextBackground, &partners); err != nil {
			return context.SendString(500, "Ошибка при загрузке партнёров")
		}

		name := context.FormValue("name")
		if name == "" {
			return context.SendString(400, "Имя партнёра обязательно")
		}

		resource, err := context.FormFile("logo_file") // ожидаем поле "logo_file"
		var logoURL string

		if err == nil {
			resourceFile, err := resource.Open()
			if err != nil {
				return context.SendString(400, "Ошибка загрузки содержимого файла логотипа")
			}
			defer resourceFile.Close()

			buf, err := io.ReadAll(resourceFile)
			if err != nil {
				return context.SendString(400, "Ошибка при чтении файла логотипа")
			}
			logoURL, err = imageUploader.UploadImage(bytes.NewReader(buf), resource.Filename)
			if err != nil {
				return context.SendString(400, "Ошибка при загрузке файла логотипа")
			}
		} else {
			logoURL = ""
		}

		newPartner := dbmodels.Partner{
			Name:      name,
			LogoURL:   logoURL,
			CreatedAt: time.Now(),
		}

		if err := db.Create(contextBackground, &newPartner); err != nil {
			return context.SendString(500, "Ошибка при создании партнёра")
		}

		if err := db.Find(contextBackground, &partners); err != nil {
			return context.SendString(500, "Ошибка при загрузке партнёров")
		}

		response := `
			<div id="partner-result" hx-swap-oob="true" style="color:green;">
				Партнёр "` + newPartner.Name + `" успешно добавлен!
			</div>

			<select id="partner-select" name="partner_id" hx-swap-oob="true">
		`

		for _, partner := range partners {
			response += `<option value="` + strconv.FormatUint(uint64(partner.ID), 10) + `">` + partner.Name + `</option>`
		}
		response += "</select>"

		return context.SendString(200, response)
	}
}
