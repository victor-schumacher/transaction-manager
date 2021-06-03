package repository

import (
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
	ID            string
	AccountID     string
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

	statement := "INSERT INTO TRANSACTION VALUES(?, ?, ?, ?, ?)"
	if _, err := db.Exec(
		statement,
		t.ID,
		t.AccountID,
		t.OperationType,
		t.Amount,
		t.EventDate,
	); err != nil {
		return err
	}

	return nil
}
