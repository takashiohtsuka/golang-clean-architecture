package router

import (
	"golang-clean-architecture/pkg/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/** ルーティング */
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ALBヘルスチェック用エンドポイント
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"status": "ok"})
	})

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	e.GET("/staffs", func(context echo.Context) error { return c.Staff.GetStaffs(context) })

	e.POST("/staffs", func(context echo.Context) error { return c.Staff.CreateStaff(context) })

	e.PUT("/staffs", func(context echo.Context) error { return c.Staff.UpdateStaff(context) })

	e.GET("/roles", func(context echo.Context) error { return c.Role.GetRoles(context) })
	e.POST("/roles", func(context echo.Context) error { return c.Role.CreateRole(context) })

	e.GET("/fanin", func(context echo.Context) error { return c.FanIn.GetFanIn(context) })
	e.GET("/urlDownloadSequential", func(context echo.Context) error {
		return c.URLDownloadSequential.GetURLDownloadSequential(context)
	})
	e.GET("/urlDownloadConcurrent", func(context echo.Context) error {
		return c.URLDownloadConcurrent.GetURLDownloadConcurrent(context)
	})

	return e
}
