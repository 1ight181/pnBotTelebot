package partner

import (
	"bytes"
	ctx "context"
	"fmt"
	"io"
	"strconv"

	adminifaces "pnBot/internal/adminpanel/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	imguploaderifaces "pnBot/internal/imageuploader/interfaces"
	loggerifaces "pnBot/internal/logger/interfaces"
)

func PartnerPost(db dbifaces.DataBaseProvider, imageUploader imguploaderifaces.ImageUploader, logger loggerifaces.Logger) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		var partners []dbmodels.Partner
		if err := db.Find(contextBackground, &partners); err != nil {
			logger.Errorf("Ошибка при загрузке партнёров: %v", err)
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при загрузке партнёров</div>")
		}

		name := context.FormValue("name")
		if name == "" {
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Имя партнёра обязательно</div>")
		}

		resource, err := context.FormFile("logo_file")
		var logoURL string

		if err == nil {
			resourceFile, err := resource.Open()
			if err != nil {
				logger.Errorf("Ошибка при открытии файла логотипа: %v", err)
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка загрузки содержимого файла логотипа</div>")
			}
			defer resourceFile.Close()

			buf, err := io.ReadAll(resourceFile)
			if err != nil {
				logger.Errorf("Ошибка при чтении файла логотипа: %v", err)
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при чтении файла логотипа</div>")
			}

			logoURL, err = imageUploader.UploadImage(bytes.NewReader(buf), resource.Filename)
			if err != nil {
				logger.Errorf("Ошибка при загрузке файла логотипа: %v", err)
				return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при загрузке файла логотипа</div>")
			}
		} else {
			logoURL = ""
		}

		newPartner := dbmodels.Partner{
			Name:    name,
			LogoUrl: logoURL,
		}

		if err := db.Create(contextBackground, &newPartner); err != nil {
			logger.Errorf("Ошибка при создании партнёра: %v", err)
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при создании партнёра</div>")
		}

		if err := db.Find(contextBackground, &partners); err != nil {
			logger.Errorf("Ошибка при загрузке партнёров после создания: %v", err)
			return context.Status(200).Type("text/html").SendString("<div class=error-box>Ошибка при загрузке партнёров</div>")
		}

		response := fmt.Sprintf(`
			<div class="success-box">Партнёр "%s" успешно добавлен!</div>
			
			<select class="input" id="partner-select" name="partner_id" hx-swap-oob="true">
		`, newPartner.Name)

		for _, partner := range partners {
			response += `<option value="` + strconv.FormatUint(uint64(partner.Id), 10) + `">` + partner.Name + `</option>`
		}
		response += "</select>"

		return context.Status(200).Type("text/html").SendString(response)
	}
}
