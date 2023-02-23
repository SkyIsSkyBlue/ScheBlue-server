package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SqlHandler struct {
}

func NewSqlHandler() SqlHandler {
	return SqlHandler{}
}

func SetRoute(e *echo.Echo, h SqlHandler) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v1")

	api.GET("/ping", h.Ping)

	return nil
}
