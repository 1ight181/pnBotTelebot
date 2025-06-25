package users

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"
	banifaces "pnBot/internal/banmanager/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
)

func UsersGet(userDao dbifaces.UserDao, banManager banifaces.BanManager) adminifaces.HandlerFunc {
	return func(context adminifaces.Context) error {
		search := context.Query("search")
		HXRequest := context.Header("HX-Request")
		isHtmx := HXRequest == "true"

		users, err := userDao.GetAllWithBans(search)
		if err != nil {
			return context.Status(500).SendString("Не удалось получить пользователей")
		}

		var view []UserWithBanInfo
		for _, user := range users {
			item := UserWithBanInfo{
				Id:       user.Id,
				TgId:     user.TgId,
				Fullname: user.Fullname,
				Username: user.Username,
			}
			isBanned, err := banManager.IsBanned(user.TgId)
			if err != nil {
				return context.Status(500).SendString("Ошибка при попытке загрузить забанненных пользователей")
			}
			if isBanned {
				item.IsBanned = true
				item.BanReason = user.BannedUser.Reason
			}
			view = append(view, item)
		}

		templateName := "userrecord"
		if !isHtmx {
			templateName = "users"
		}

		return context.Render(200, templateName, map[string]any{
			"Users": view,
		})
	}
}
