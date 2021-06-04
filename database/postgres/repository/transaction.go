package repository

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
	"transaction-manager/database"
)

type Transaction interface {
	Save(transaction TransactionEntity) error
}

type TransactionRepo struct {
	db database.DBConnection
}

type TransactionEntity struct {
	ID            uuid.UUID
	AccountID     uuid.UUID
	OperationType int
	Amount        int
	EventDate     time.Time
}

func NewTransaction(db database.DBConnection) TransactionRepo {
	return TransactionRepo{db: db}
}

func (tr TransactionRepo) Save(t TransactionEntity) error {
	db := tr.db.Connect()
	defer db.Close()

	statement := "INSERT INTO TRANSACTION VALUES($1, $2, $3, $4, $5)"
	if _, err := db.Exec(
		statement,
		t.ID,
		t.AccountID,
		t.OperationType,
		t.Amount,
		t.EventDate,
	); err != nil {
		log.Err(err).Msg("cannot add transaction into database")
		return err
	}

	return nil
}
