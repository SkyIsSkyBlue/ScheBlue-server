package router

import (
	"log"
	"net/http"

	"github.com/SkyIsSkyBlue/ScheBlue-server/model"
	"github.com/labstack/echo/v4"
)

func (h *SqlHandler) Ping(ctx echo.Context) error {
	log.Println("ping recieved")

	pong := "pong"

	return ctx.String(http.StatusOK, pong)
}

func (h *SqlHandler) SqlPing(ctx echo.Context) error {
	log.Println("sqlping recieved")

	var sqlping model.SqlPing

	err := h.DB.Where(&model.SqlPing{PingId: "sqlPing"}).Find(&sqlping).Error
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, sqlping.PongValue)
}
