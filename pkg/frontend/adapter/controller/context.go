package controller

import "net/http"

type Context interface {
	JSON(code int, i any) error
	Bind(i any) error
	Param(name string) string
	QueryParam(name string) string
	Request() *http.Request
}
