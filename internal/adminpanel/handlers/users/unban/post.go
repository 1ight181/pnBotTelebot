package delete

import (
	"fmt"
	"strconv"

	adminifaces "pnBot/internal/adminpanel/interfaces"
	banifaces "pnBot/internal/banmanager/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
)

func UnbanUserPost(banManager banifaces.BanManager, userDao dbifaces.UserDao) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ids := context.FormValues("user_ids")
		var unbannedCount int

		for _, idStr := range ids {
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				return context.Status(200).Type("text/html").
					SendString(`<div id="user-action-feedback" class="error-box" hx-swap-oob="true">Ошибка при чтении ID пользователя</div> 
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			tgID, err := userDao.GetTgIdByUserId(uint(id))
			if err != nil {
				return context.Status(200).Type("text/html").
					SendString(`<div id="user-action-feedback" class="error-box" hx-swap-oob="true">Пользователь не найден</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			if err := banManager.Unban(tgID); err != nil {
				return context.Status(200).Type("text/html").
					SendString(`<div id="user-action-feedback" class="error-box" hx-swap-oob="true">Ошибка при разбане</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			unbannedCount++
		}

		if unbannedCount == 0 {
			return context.Status(200).Type("text/html").
				SendString(`<div id="user-action-feedback" class="error-box" hx-swap-oob="true">Ни одного пользователя не разбанено</div>
				<div
					hx-get="/users"
					hx-trigger="load"
					hx-swap="innerHTML"
				>
					{{ template "userrecord" . }}
				</div>`)
		}

		return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`
			<div id="user-action-feedback" class="success-box" hx-swap-oob="true">
				Пользователи в количестве %d успешно разбанены
			</div>
			<div
				hx-get="/users"
				hx-trigger="load"
				hx-swap="innerHTML"
			>
				{{ template "userrecord" . }}
			</div>
		`, unbannedCount))
	}
}
