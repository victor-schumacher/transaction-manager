package main

import (
	"github.com/labstack/echo"
	"log"
	"transaction-manager/config"
	"transaction-manager/database/postgres"
	"transaction-manager/http/handler"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()
	ah := handler.NewAccount()
	th := handler.NewTransaction()
	db := postgres.NewConnection(c)
	db.TestConnection()

	ah.Handle(e)
	th.Handle(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
