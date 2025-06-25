package interfaces

type SpamManager interface {
	IsAllowed(userId int64) (bool, bool, int, error)
}
