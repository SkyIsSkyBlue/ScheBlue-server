package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *SqlHandler) Ping(ctx echo.Context) error {
	log.Println("ping recieved")

	pong := "pong"

	return ctx.String(http.StatusOK, pong)
}
