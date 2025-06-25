package ban

import (
	"fmt"
	"strconv"
	"time"

	adminifaces "pnBot/internal/adminpanel/interfaces"
	banifaces "pnBot/internal/banmanager/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
)

func BanUserPost(banManager banifaces.BanManager, userDao dbifaces.UserDao) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		ids := context.FormValues("user_ids")
		var bannedCount int

		for _, idStr := range ids {
			id, _ := strconv.ParseUint(idStr, 10, 64)
			tgId, err := userDao.GetTgIdByUserId(uint(id))
			if err != nil {
				return context.Status(200).Type("text/html").SendString(`<div id="user-action-feedback" hx-swap-oob="true" class="error-box">Ошибка при получении TgId пользователя</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			isBanned, err := banManager.IsBanned(tgId)
			if err != nil {
				return context.Status(200).Type("text/html").SendString(`<div id="user-action-feedback" hx-swap-oob="true" class="error-box">Ошибка при попытке узнать забанен ли пользователь</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			if isBanned {
				return context.Status(200).Type("text/html").SendString(`<div id="user-action-feedback" hx-swap-oob="true" class="error-box">Ошибка: один из выбранных пользователей уже забанен</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			err = banManager.Ban(tgId, "Manual ban", time.Duration(0), "admin panel")
			if err != nil {
				return context.Status(200).Type("text/html").SendString(`<div id="user-action-feedback" hx-swap-oob="true" class="error-box">Ошибка при бане пользователя</div>
					<div
						hx-get="/users"
						hx-trigger="load"
						hx-swap="innerHTML"
					>
						{{ template "userrecord" . }}
					</div>`)
			}
			bannedCount++
		}

		if bannedCount == 0 {
			return context.Status(200).Type("text/html").SendString(`<div id="user-action-feedback" hx-swap-oob="true" class="error-box">Не удалось забанить ни одного пользователя</div>
				<div
					hx-get="/users"
					hx-trigger="load"
					hx-swap="innerHTML"
				>
					{{ template "userrecord" . }}
				</div>`)
		}

		return context.Status(200).Type("text/html").SendString(fmt.Sprintf(`
			<div id="user-action-feedback" hx-swap-oob="true" class="success-box">
				Пользователи в количестве %d успешно забанены!
			</div>
			<div
				hx-get="/users"
				hx-trigger="load"
				hx-swap="innerHTML"
			>
				{{ template "userrecord" . }}
			</div>
		`, bannedCount))
	}
}
