package controller

import "net/http"

type Context interface {
	JSON(code int, i any) error
	Bind(i any) error
	Request() *http.Request
}
