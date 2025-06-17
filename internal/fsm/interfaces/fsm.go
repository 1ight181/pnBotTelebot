package interfaces

type Fsm interface {
	Set(userID int64, state string)
	Get(userID int64) string
	Clear(userID int64)
}
