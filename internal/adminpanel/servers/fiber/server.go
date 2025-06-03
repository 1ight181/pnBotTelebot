package fiber

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

func New() *FiberServer {
	return &FiberServer{app: fiber.New()}
}

func (s *FiberServer) GET(path string, handler adminifaces.HandlerFunc) {
	s.app.Get(path, func(c *fiber.Ctx) error {
		return handler(&fiberContext{c})
	})
}

func (s *FiberServer) POST(path string, handler adminifaces.HandlerFunc) {
	s.app.Post(path, func(c *fiber.Ctx) error {
		return handler(&fiberContext{c})
	})
}

func (s *FiberServer) Listen(addr string) error {
	return s.app.Listen(addr)
}
