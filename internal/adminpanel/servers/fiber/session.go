package fiber

import "github.com/gofiber/fiber/v2/middleware/session"

type FiberSession struct {
	session *session.Session
}

func NewFiberSession(sess *session.Session) *FiberSession {
	return &FiberSession{session: sess}
}

func (fs *FiberSession) Get(key string) interface{} {
	return fs.session.Get(key)
}

func (fs *FiberSession) Set(key string, value interface{}) {
	fs.session.Set(key, value)
}

func (fs *FiberSession) Delete(key string) {
	fs.session.Delete(key)
}

func (fs *FiberSession) Save() error {
	return fs.session.Save()
}

func (fs *FiberSession) Destroy() error {
	return fs.session.Destroy()
}
