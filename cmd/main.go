package main

import (
	"github.com/labstack/echo"
	"transaction-manager/http/handler"
)

func main() {
	e := echo.New()
	h := handler.NewAccount()
	h.MakeAccountHandlers(e)
	if err := e.Start(":8080"); err != nil {
		return
	}
}
