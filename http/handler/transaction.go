package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"transaction-manager/database/postgres"
	"transaction-manager/database/postgres/repository"
)

type Transaction struct {
	AccountID       uuid.UUID `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          int       `json:"amount"`
	Balance         int       `json:"balance"`
}

type TransactionResponse struct {
	Message string `json:"message"`
}

type TransactionManager struct {
	repo repository.Transaction
	e    *echo.Echo
}

func NewTransaction(
	repo repository.Transaction,
	e *echo.Echo,
) TransactionManager {
	return TransactionManager{
		repo: repo,
		e:    e,
	}
}

func (m TransactionManager) createNew(c echo.Context) error {
	t := Transaction{}
	if err := c.Bind(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	te := repository.TransactionEntity{
		ID:            uuid.New(),
		AccountID:     t.AccountID,
		OperationType: t.OperationTypeID,
		Amount:        t.Amount,
		Balance:       t.Balance,
		EventDate:     time.Now(),
	}

	err := m.repo.Save(te)
	if err == postgres.ErrNotAvailableCreditLimit {
		return echo.NewHTTPError(http.StatusBadRequest, "no available limit")
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, TransactionResponse{"transaction successful"})
}
func (m TransactionManager) Handle() {
	m.e.POST("/transactions", m.createNew)
}
