package banmanager

import (
	ctx "context"
	"fmt"
	cacheifaces "pnBot/internal/cache/interfaces"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"time"
)

type BanManager struct {
	dbProvider    dbifaces.DataBaseProvider
	context       ctx.Context
	cacheProvider cacheifaces.CacheProvider
}

func NewBanManager(context ctx.Context, dbProvider dbifaces.DataBaseProvider, cacheProvider cacheifaces.CacheProvider) *BanManager {
	return &BanManager{
		dbProvider:    dbProvider,
		context:       context,
		cacheProvider: cacheProvider,
	}
}

func (bm *BanManager) Ban(userId int64, reason string, duration time.Duration, author string) error {
	isBannedKey := fmt.Sprintf("user:ban:%d", userId)

	var user dbmodels.User
	err := bm.dbProvider.First(bm.context, &user, "tg_id = ?", userId)
	if err != nil {
		return err
	}

	var existing dbmodels.BannedUser
	err = bm.dbProvider.First(bm.context, &existing, "user_id = ?", user.Id)
	if err == nil {
		if existing.ExpiresAt == nil || existing.ExpiresAt.After(time.Now()) {
			return nil
		}
	}

	exp := time.Now().Add(duration)
	ban := &dbmodels.BannedUser{
		UserId:    user.Id,
		Reason:    reason,
		ExpiresAt: &exp,
		CreatedBy: author,
	}

	if err := bm.cacheProvider.Set(isBannedKey, "true", duration); err != nil {
		return err
	}

	return bm.dbProvider.Create(bm.context, ban)
}

func (bm *BanManager) IsBanned(userId int64) (bool, error) {
	var user dbmodels.User
	if err := bm.dbProvider.First(bm.context, &user, "tg_id = ?", userId); err != nil {
		return false, err
	}

	var ban dbmodels.BannedUser
	if err := bm.dbProvider.First(bm.context, &ban, "user_id = ?", user.Id); err != nil {
		return false, nil
	}

	if ban.ExpiresAt == nil || ban.ExpiresAt.After(time.Now()) {
		return true, nil
	}
	return false, nil
}

func (bm *BanManager) Unban(userId int64) error {
	isBannedKey := fmt.Sprintf("user:ban:%d", userId)
	warnCountKey := fmt.Sprintf("user:warncount:%d", userId)
	messageCountKey := fmt.Sprintf("user:msgcount:%d", userId)

	var user dbmodels.User
	if err := bm.dbProvider.First(bm.context, &user, "tg_id = ?", userId); err != nil {
		return err
	}

	where := dbmodels.BannedUser{
		UserId: user.Id,
	}

	if err := bm.cacheProvider.Del(isBannedKey); err != nil {
		return err
	}

	if err := bm.cacheProvider.Del(warnCountKey); err != nil {
		return err
	}

	if err := bm.cacheProvider.Del(messageCountKey); err != nil {
		return err
	}

	return bm.dbProvider.Delete(
		bm.context,
		&dbmodels.BannedUser{},
		where,
	)
}
