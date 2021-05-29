package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type Account struct {
	ID string `json:"account_id"`
}

func NewAccount() Account {
	return Account{}
}

func (a Account) findById(c echo.Context) error {
	return c.JSON(http.StatusOK, Account{ID: "1"})
}

func (a Account) createNew(c echo.Context) error {
	return c.JSON(http.StatusCreated, Account{ID: "1"})
}

func (a Account) Handle(e *echo.Echo) {
	e.GET("/accounts/:id", a.findById)
	e.POST("/accounts", a.createNew)
}
