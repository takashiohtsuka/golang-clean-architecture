package router

import (
	"golang-clean-architecture/pkg/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/** ルーティング */
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	e.GET("/staffs", func(context echo.Context) error { return c.Staff.GetStaffs(context) })

	e.POST("/staffs", func(context echo.Context) error { return c.Staff.CreateStaff(context) })

	e.PUT("/staffs", func(context echo.Context) error { return c.Staff.UpdateStaff(context) })

	e.POST("/roles", func(context echo.Context) error { return c.Role.CreateRole(context) })

	return e
}
