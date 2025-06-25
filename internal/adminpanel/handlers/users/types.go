package users

type UserWithBanInfo struct {
	Id        uint
	TgId      int64
	Fullname  string
	Username  string
	IsBanned  bool
	BanReason string
}
