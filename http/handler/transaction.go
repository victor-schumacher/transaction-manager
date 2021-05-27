package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type Transaction struct {
	AccountID       string `json:"account_id"`
	OperationTypeID int    `json:"operation_type_id"`
	Amount          int    `json:"amount"`
}

func NewTransaction() Transaction {
	return Transaction{}
}

func (a Transaction) createNew(c echo.Context) error {
	t := Transaction{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, t)
}

func (a Transaction) Handle(e *echo.Echo) {
	e.POST("/transactions", a.createNew)
}
