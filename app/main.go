package main

import (
	"github.com/SkyIsSkyBlue/ScheBlue-server/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	h, err := router.NewSqlHandler()
	if err != nil {
		panic(err)
	}

	err = router.SetRoute(e, h)
	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
