package fiber

import (
	"mime/multipart"
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2"
)

type FiberContext struct {
	context *fiber.Ctx
}

func (fc *FiberContext) Param(name string) string {
	return fc.context.Params(name)
}

func (fc *FiberContext) Query(name string) string {
	return fc.context.Query(name)
}

func (fc *FiberContext) Header(name string) string {
	return fc.context.Get(name)
}

func (fc *FiberContext) Method() string {
	return fc.context.Method()
}

func (fc *FiberContext) Path() string {
	return fc.context.Path()
}

func (fc *FiberContext) BodyParser(out interface{}) error {
	return fc.context.BodyParser(out)
}

func (fc *FiberContext) Cookie(name string) string {
	return fc.context.Cookies(name)
}

func (fc *FiberContext) JSON(code int, data interface{}) error {
	return fc.context.Status(code).JSON(data)
}

func (fc *FiberContext) SendString(data string) error {
	return fc.context.SendString(data)
}

func (fc *FiberContext) Status(code int) adminifaces.Context {
	fc.context.Status(code)
	return fc
}

func (fc *FiberContext) Type(contentType string) adminifaces.Context {
	fc.context.Type(contentType)
	return fc
}

func (fc *FiberContext) FormFile(name string) (*multipart.FileHeader, error) {
	return fc.context.FormFile(name)
}

func (fc *FiberContext) FormValue(name string) string {
	return fc.context.FormValue(name)
}

func (fc *FiberContext) SetHeader(key, value string) {
	fc.context.Set(key, value)
}

func (fc *FiberContext) SetCookie(name, value string, maxAge int) {
	fc.context.Cookie(&fiber.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: maxAge,
	})
}

func (fc *FiberContext) Render(code int, name string, data map[string]any) error {
	return fc.context.Status(code).Render(name, data)
}

func (fc *FiberContext) Redirect(location string, status ...int) error {
	if len(status) > 0 {
		return fc.context.Redirect(location, status[0])
	}
	return fc.context.Redirect(location)
}

func (fc *FiberContext) Next() error {
	return fc.context.Next()
}

func (fc *FiberContext) Context() any {
	return fc.context
}
