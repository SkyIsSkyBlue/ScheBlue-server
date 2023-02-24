package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo, h SqlHandler) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v1")
	{
		api.GET("/ping", h.Ping)
		api.GET("/sqlping", h.SqlPing)
	}

	return nil
}
