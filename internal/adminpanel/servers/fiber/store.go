package fiber

import (
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type FiberSessionStore struct {
	store *session.Store
}

func NewSessionStore(store *session.Store) *FiberSessionStore {
	return &FiberSessionStore{
		store: store,
	}
}

func (s *FiberSessionStore) Get(ctx adminifaces.Context) (*session.Session, error) {
	fiberContext, ok := ctx.Context().(*fiber.Ctx)
	if !ok {
		return nil, fiberContext.Status(500).SendString("invalid context")
	}

	session, err := s.store.Get(fiberContext)
	if err != nil {
		return nil, err
	}

	return session, nil
}
