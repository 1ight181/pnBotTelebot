package interfaces

type Session interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
	Save() error
	Destroy() error
}
