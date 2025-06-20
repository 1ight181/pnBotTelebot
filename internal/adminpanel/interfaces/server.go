package interfaces

type HandlerFunc func(Context) error

type Server interface {
	Use(path string, middleware HandlerFunc)
	Static(prefix string, root string)
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	Listen(addr string) error
	Shutdown() error
}
