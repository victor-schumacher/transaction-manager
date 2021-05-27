package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type Account struct {
}

func NewAccount() Account {
	return Account{}
}

func (a Account) findById(c echo.Context) error {
	return c.JSON(http.StatusOK, `{"account_id":1, "document_number":"09365523859"}`)
}

func (a Account) createNew(c echo.Context) error {
	return c.JSON(http.StatusCreated, `{"account_id":"1"}`)
}

func (a Account) Handle(e *echo.Echo) {
	e.GET("/accounts", a.findById)
	e.POST("accounts/", a.createNew)
}
