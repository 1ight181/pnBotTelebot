package fiber

import (
	"fmt"
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app *fiber.App
}

func NewFiberServer(app *fiber.App) *FiberServer {
	return &FiberServer{app: app}
}

func (fs *FiberServer) Use(path string, middleware adminifaces.HandlerFunc) {
	fs.app.Use(path, func(c *fiber.Ctx) error {
		fiberContext := &FiberContext{context: c}
		return middleware(fiberContext)
	})
}

func (fs *FiberServer) Static(prefix string, root string) {
	fs.app.Static(prefix, root)
}

func (fs *FiberServer) GET(path string, handler adminifaces.HandlerFunc) {
	fs.app.Get(path, func(c *fiber.Ctx) error {
		fiberContext := &FiberContext{context: c}
		return handler(fiberContext)
	})
}

func (fs *FiberServer) POST(path string, handler adminifaces.HandlerFunc) {
	fs.app.Post(path, func(c *fiber.Ctx) error {
		fiberContext := &FiberContext{context: c}
		return handler(fiberContext)
	})
}

func (fs *FiberServer) Listen(addr string) error {
	return fs.app.Listen(addr)
}

func (fs *FiberServer) Shutdown() error {
	if err := fs.app.Shutdown(); err != nil {
		return fmt.Errorf("не удалось корректно завершить FiberServer: %v", err)
	}
	return nil
}
