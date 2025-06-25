package delete

import (
	"fmt"
	"strconv"

	adminifaces "pnBot/internal/adminpanel/interfaces"
	banifaces "pnBot/internal/banmanager/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
)

func DeleteUserPost(userDao dbifaces.UserDao, banManager banifaces.BanManager) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ids := context.FormValues("user_ids")
		var deletedCount int

		for _, idStr := range ids {
			id, _ := strconv.ParseUint(idStr, 10, 64)

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

			if err := userDao.Delete(uint(id)); err != nil {
				return context.Status(200).Type("text/html").
					SendString(`<div id="user-action-feedback" class="error-box" hx-swap-oob="true">Ошибка при удалении пользователя</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}

			deletedCount++
		}

		if deletedCount == 0 {
			return context.Status(200).Type("text/html").
				SendString(`<div id="user-action-feedback" class="error-box" hx-swap-oob="true">Не удалось удалить пользователей</div><div
				hx-get="/users"
				hx-trigger="load"
				hx-swap="innerHTML"
			>
				{{ template "userrecord" . }}
			</div>`)
		}

		return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`
			<div id="user-action-feedback" class="success-box" hx-swap-oob="true">
				Пользователи в количестве %d успешно удалены
			</div>
			<div
				hx-get="/users"
				hx-trigger="load"
				hx-swap="innerHTML"
			>
				{{ template "userrecord" . }}
			</div>
		`, deletedCount))
	}
}
