package main

import (
	"github.com/labstack/echo"
	"log"
	"transaction-manager/config"
	"transaction-manager/database/postgres"
	"transaction-manager/database/postgres/repository"
	"transaction-manager/http/handler"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()
	db := postgres.NewConnection(c)
	db.TestConnection()

	accountRepo := repository.NewAccount(db)
	transactionRepo := repository.NewTransaction(db)

	ah := handler.NewAccount(accountRepo, e)
	th := handler.NewTransaction(transactionRepo, e)

	ah.Handle()
	th.Handle()

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
