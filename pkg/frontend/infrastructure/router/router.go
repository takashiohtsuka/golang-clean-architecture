package router

import (
	"golang-clean-architecture/pkg/frontend/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	g := e.Group("/frontend", middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET},
	}))

	g.GET("/stores", func(ctx echo.Context) error { return c.Store.GetStoreList(ctx) })
	g.GET("/stores/:id", func(ctx echo.Context) error { return c.Store.GetStoreDetail(ctx) })

	g.GET("/women", func(ctx echo.Context) error { return c.Woman.GetWomanList(ctx) })
	g.GET("/women/:id", func(ctx echo.Context) error { return c.Woman.GetWomanDetail(ctx) })

	return e
}
