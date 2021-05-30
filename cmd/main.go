package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"transaction-manager/config"
	"transaction-manager/http/handler"
)

func main() {
	e := echo.New()
	ah := handler.NewAccount()
	ah.Handle(e)
	th := handler.NewTransaction()
	th.Handle(e)

	c, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(c)
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
