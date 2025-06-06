package interfaces

type SessionStore[session interface{}] interface {
	Get(ctx Context) (session, error)
}
