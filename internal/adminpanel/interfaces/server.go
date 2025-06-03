package interfaces

type HandlerFunc func(Context) error

type Server interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	Listen(addr string) error
}
