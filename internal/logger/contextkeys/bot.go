package contextkeys

type contextKey string

const (
	UserIDKey contextKey = "userID"
	ChatIDKey contextKey = "chatID"
	TextKey   contextKey = "text"
	DataKey   contextKey = "data"
)
