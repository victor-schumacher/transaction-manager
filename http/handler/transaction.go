package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"transaction-manager/database/postgres/repository"
)

type Transaction struct {
	AccountID       uuid.UUID `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          int       `json:"amount"`
}

type TransactionManager struct {
	repo repository.TransactionRepo
	e    *echo.Echo
}

func NewTransaction(
	repo repository.TransactionRepo,
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
		return err
	}
	te := repository.TransactionEntity{
		ID:            uuid.New(),
		AccountID:     t.AccountID,
		OperationType: t.OperationTypeID,
		Amount:        t.Amount,
		EventDate:     time.Now(),
	}

	if err := m.repo.Save(te); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

func (m TransactionManager) Handle() {
	m.e.POST("/transactions", m.createNew)
}
