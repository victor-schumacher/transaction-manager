package handler

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
	"transaction-manager/database/postgres/repository"
)

type AccountManager struct {
	repo repository.Account
	e    *echo.Echo
}

type Account struct {
	ID             uuid.UUID `json:"account_id"`
	DocumentNumber string    `json:"document_number"`
}

type AccountCreatedResponse struct {
	Message string `json:"message"`
}

func NewAccount(
	repo repository.Account,
	e *echo.Echo,
) AccountManager {
	return AccountManager{
		repo: repo,
		e:    e,
	}
}

func (m AccountManager) findById(c echo.Context) error {
	paramId := c.Param("id")
	id, err := uuid.Parse(paramId)
	if err != nil {
		log.Err(err).Msg("cannot parse uuid at find one account endpoint")
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid id")
	}

	accountEntity, err := m.repo.FindOne(id)
	if err == sql.ErrNoRows {
		return echo.NewHTTPError(
			http.StatusNotFound,
			"account not found!",
		)
	}
	if err != nil {
		log.Err(err).Msg("cannot parse uuid at find one account endpoint")
		return echo.NewHTTPError(http.StatusInternalServerError, nil)
	}

	account := Account{
		ID:             accountEntity.ID,
		DocumentNumber: accountEntity.DocumentNumber,
	}
	return c.JSON(http.StatusOK, account)
}

func (m AccountManager) createNew(c echo.Context) error {
	a := Account{}
	if err := c.Bind(&a); err != nil {
		return err
	}

	ae := repository.AccountEntity{
		ID:             uuid.New(),
		DocumentNumber: a.DocumentNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := m.repo.Save(ae); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, AccountCreatedResponse{
		Message: "account successfully created"},
	)
}

func (m AccountManager) Handle() {
	m.e.GET("/accounts/:id", m.findById)
	m.e.POST("/accounts", m.createNew)
}
