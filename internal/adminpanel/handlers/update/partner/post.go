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

func UpdatePartnerPost(
	db dbifaces.DataBaseProvider,
	imageUploader imguploaderifaces.ImageUploader,
	logger loggerifaces.Logger,
) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		partnerID, err := strconv.Atoi(context.FormValue("id"))
		if err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Некорректный ID партнёра</div>")
		}

		var partner dbmodels.Partner
		if err := db.First(contextBackground, &partner, "id = ?", partnerID); err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Партнёр не найден</div>")
		}

		name := context.FormValue("name")
		if name == "" {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Название партнёра обязательно</div>")
		}

		resource, err := context.FormFile("logo_file")
		if err == nil {
			resourceFile, err := resource.Open()
			if err != nil {
				logger.Errorf("Ошибка при открытии файла логотипа: %v", err)
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка загрузки логотипа</div>")
			}
			defer resourceFile.Close()

			buf, err := io.ReadAll(resourceFile)
			if err != nil {
				logger.Errorf("Ошибка при чтении файла логотипа: %v", err)
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка при чтении логотипа</div>")
			}

			logoURL, err := imageUploader.UploadImage(bytes.NewReader(buf), resource.Filename)
			if err != nil {
				logger.Errorf("Ошибка при загрузке логотипа: %v", err)
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка при загрузке логотипа</div>")
			}
			partner.LogoUrl = logoURL
		}

		partner.Name = name

		if err := db.Save(contextBackground, &partner); err != nil {
			logger.Errorf("Ошибка при обновлении партнёра: %v", err)
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Ошибка при сохранении партнёра</div>")
		}

		var partners []dbmodels.Partner
		if err := db.Find(contextBackground, &partners); err != nil {
			logger.Errorf("Ошибка при повторной загрузке партнёров: %v", err)
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Ошибка при загрузке партнёров</div>")
		}

		response := `
			<div class="success-box">Партнёр успешно обновлён!</div>
			<select id="partner-select" name="partner_id" hx-swap-oob="true"
				hx-get="/update/partners"
				hx-target="#partner-edit-form"
				hx-swap="innerHTML"
				class="input"
			>`

		response += `<option value="">-- Выберите партнёра --</option>`
		for _, partner := range partners {
			selected := ""
			if partner.Id == uint(partnerID) {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, partner.Id, selected, partner.Name)
		}
		response += "</select>"

		return context.Status(200).Type("text/html").SendString(response)
	}
}
