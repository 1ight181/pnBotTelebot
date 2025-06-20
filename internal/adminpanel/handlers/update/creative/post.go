package creative

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
	imageutils "pnBot/internal/imageutils"
)

func UpdateCreativePost(
	db dbifaces.DataBaseProvider,
	imageUploader imguploaderifaces.ImageUploader,
) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		contextBackground := ctx.Background()

		creativeIDStr := context.FormValue("id")
		creativeID, err := strconv.Atoi(creativeIDStr)
		if err != nil {
			return context.Type("text/html").Status(200).
				SendString("<div class=error-box>Некорректный ID креатива</div>")
		}

		var creative dbmodels.Creative
		if err := db.First(contextBackground, &creative, "id = ?", creativeID); err != nil {
			return context.Type("text/html").Status(200).
				SendString("<div class=error-box>Креатив не найден</div>")
		}

		partnerID := context.FormValue("partner_internal_creative_id")
		typeStr := context.FormValue("type")
		offerIDStr := context.FormValue("offer_id")

		if partnerID == "" || offerIDStr == "" {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Обязательные поля не заполнены</div>")
		}

		offerID, err := strconv.ParseUint(offerIDStr, 10, 64)
		if err != nil {
			return context.Status(200).Type("text/html").
				SendString("<div class=error-box>Некорректный ID оффера</div>")
		}

		creative.PartnerInternalCreativeId = partnerID
		creative.Type = typeStr
		creative.OfferId = uint(offerID)

		resource, err := context.FormFile("image")
		if err == nil {
			file, err := resource.Open()
			if err != nil {
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка открытия файла изображения</div>")
			}
			defer file.Close()

			buf, err := io.ReadAll(file)
			if err != nil {
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка чтения изображения</div>")
			}

			imageURL, err := imageUploader.UploadImage(bytes.NewReader(buf), resource.Filename)
			if err != nil {
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка загрузки изображения</div>")
			}

			width, height, err := imageutils.GetImageDimensions(bytes.NewReader(buf))
			if err != nil {
				return context.Status(200).Type("text/html").
					SendString("<div class=error-box>Ошибка определения размеров изображения</div>")
			}

			creative.ResourceUrl = imageURL
			creative.Width = width
			creative.Height = height
		}

		if err := db.Save(contextBackground, &creative); err != nil {
			return context.Type("text/html").Status(200).
				SendString("<div class=error-box>Ошибка при обновлении креатива</div>")
		}

		var creatives []dbmodels.Creative
		if err := db.Find(contextBackground, &creatives); err != nil {
			return context.Type("text/html").Status(200).SendString("<div class=error-box>Ошибка при загрузке креативов</div>")
		}

		response := `
			<div class="success-box">Креатив успешно обновлен!</div>

			<select 
				id="creative-select" 
				hx-swap-oob="true" 
				class="input"
				hx-get="/update/creatives"
				hx-target="#creative-edit-form"
				hx-swap="innerHTML"
				name="creative_id"
			>
		`

		response += `<option value="">-- Выберите креатив для редактирования --</option>`

		for _, creative := range creatives {
			selected := ""
			if creative.Id == uint(creativeID) {
				selected = "selected"
			}
			response += fmt.Sprintf(`<option value="%d" %s>%s</option>`, creative.Id, selected, creative.PartnerInternalCreativeId)
		}
		response += "</select>"
		return context.Status(200).Type("text/html").SendString(response)
	}
}
