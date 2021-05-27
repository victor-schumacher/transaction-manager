package main

import (
	"github.com/labstack/echo"
	"transaction-manager/http/handler"
)

func main() {
	e := echo.New()
	ah := handler.NewAccount()
	ah.Handle(e)
	th := handler.NewTransaction()
	th.Handle(e)
	if err := e.Start(":8080"); err != nil {
		return
	}
}
