package banmanager

import (
	ctx "context"
	dbifaces "pnBot/internal/db/interfaces"
	dbmodels "pnBot/internal/db/models"
	"time"
)

type BanManager struct {
	dbProvider dbifaces.DataBaseProvider
	context    ctx.Context
}

func NewBanManager(context ctx.Context, dbProvider dbifaces.DataBaseProvider) *BanManager {
	return &BanManager{
		dbProvider: dbProvider,
		context:    context,
	}
}

func (bm *BanManager) Ban(userId int64, reason string, duration time.Duration, author string) error {
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
		UserID:    user.Id,
		Reason:    reason,
		ExpiresAt: &exp,
		CreatedBy: author,
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
	var user dbmodels.User
	if err := bm.dbProvider.First(bm.context, &user, "tg_id = ?", userId); err != nil {
		return err
	}

	where := dbmodels.BannedUser{
		UserID: user.Id,
	}
	expiredTime := time.Now().Add(-time.Second)
	return bm.dbProvider.Update(bm.context, where, "expires_at", &expiredTime)
}
