package interfaces

import "mime/multipart"

type Context interface {
	Param(name string) string
	Query(name string) string
	Header(name string) string
	Method() string
	Path() string
	BodyParser(out interface{}) error
	Cookie(name string) string
	FormValue(name string) string
	FormFile(name string) (*multipart.FileHeader, error)

	JSON(code int, data interface{}) error
	SendString(code int, data string) error
	Status(code int) Context
	SetHeader(key, value string)
	SetCookie(name, value string, maxAge int)
	Render(code int, name string, data map[string]any) error
	Redirect(location string, status ...int) error

	Context() any
}
