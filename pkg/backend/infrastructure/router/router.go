package router

import (
	"golang-clean-architecture/pkg/backend/adapter/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	g := e.Group("/backend")

	g.POST("/companies", func(ctx echo.Context) error { return c.Company.CreateCompany(ctx) })
	g.POST("/stores", func(ctx echo.Context) error { return c.Store.CreateStore(ctx) })
	g.PUT("/stores/:id", func(ctx echo.Context) error { return c.Store.UpdateStore(ctx) })
	g.POST("/management_staffs", func(ctx echo.Context) error { return c.ManagementStaff.CreateManagementStaff(ctx) })
	g.POST("/women", func(ctx echo.Context) error { return c.Woman.CreateWoman(ctx) })
	g.PUT("/women/:id", func(ctx echo.Context) error { return c.Woman.UpdateWoman(ctx) })

	return e
}
