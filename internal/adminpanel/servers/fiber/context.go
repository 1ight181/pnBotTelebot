package fiber

import (
	"mime/multipart"
	adminifaces "pnBot/internal/adminpanel/interfaces"

	"github.com/gofiber/fiber/v2"
)

type fiberContext struct {
	context *fiber.Ctx
}

func (fc *fiberContext) Param(name string) string {
	return fc.context.Params(name)
}

func (fc *fiberContext) Query(name string) string {
	return fc.context.Query(name)
}

func (fc *fiberContext) Header(name string) string {
	return fc.context.Get(name)
}

func (fc *fiberContext) Method() string {
	return fc.context.Method()
}

func (fc *fiberContext) Path() string {
	return fc.context.Path()
}

func (fc *fiberContext) BodyParser(out interface{}) error {
	return fc.context.BodyParser(out)
}

func (fc *fiberContext) Cookie(name string) string {
	return fc.context.Cookies(name)
}

func (fc *fiberContext) JSON(code int, data interface{}) error {
	return fc.context.Status(code).JSON(data)
}

func (fc *fiberContext) SendString(code int, data string) error {
	return fc.context.Status(code).SendString(data)
}

func (fc *fiberContext) Status(code int) adminifaces.Context {
	fc.context.Status(code)
	return fc
}

func (fc *fiberContext) FormFile(name string) (*multipart.FileHeader, error) {
	return fc.context.FormFile(name)
}

func (fc *fiberContext) FormValue(name string) string {
	return fc.context.FormValue(name)
}

func (fc *fiberContext) SetHeader(key, value string) {
	fc.context.Set(key, value)
}

func (fc *fiberContext) SetCookie(name, value string, maxAge int) {
	fc.context.Cookie(&fiber.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: maxAge,
	})
}

func (fc *fiberContext) Render(code int, name string, data map[string]any) error {
	return fc.context.Status(code).Render(name, data)
}

func (fc *fiberContext) Redirect(location string, status ...int) error {
	if len(status) > 0 {
		return fc.context.Redirect(location, status[0])
	}
	return fc.context.Redirect(location)
}

func (fc *fiberContext) Context() any {
	return fc.context
}
