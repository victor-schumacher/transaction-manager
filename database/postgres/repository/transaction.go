package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
	"transaction-manager/database"
	"transaction-manager/database/postgres"
)

const (
	payment = 4
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
	Balance       int
	EventDate     time.Time
}

func NewTransaction(db database.DBConnection) TransactionRepo {
	return TransactionRepo{db: db}
}

func (tr TransactionRepo) Save(t TransactionEntity) error {
	db := tr.db.Connect()
	defer db.Close()
	availableCreditLimit := 0
	row := db.QueryRow("SELECT available_credit_limit from account WHERE id=$1", t.AccountID)
	if err := row.Scan(&availableCreditLimit); err != nil {
		log.Err(err).Msg("cannot get available credit limit from database")
		return err
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Err(err)
		return err
	}
	newLimit := availableCreditLimit - t.Amount
	if t.OperationType == payment {
		newLimit = availableCreditLimit + t.Amount

		rows, err := tx.Query(`SELECT balance FROM transaction.transaction  WHERE transaction.transaction.account_id=$1  ORDER  BY event_date DESC`, t.AccountID)
		if err != nil {
			tx.Rollback()
			log.Err(err)
		}

		var amounts []int
		for rows.Next() {
			amount := 0
			if err := rows.Scan(&amount); err != nil {
				tx.Rollback()
				log.Err(err)
			}
			amounts = append(amounts, amount)
		}

				

	}

	if t.Amount > availableCreditLimit && t.OperationType != payment {
		log.Err(postgres.ErrNotAvailableCreditLimit).Msg("limit exceed")
		return postgres.ErrNotAvailableCreditLimit
	}

	_, err = tx.ExecContext(ctx, "UPDATE account SET available_credit_limit=$1 WHERE ID=$2;", newLimit, t.AccountID)
	if err != nil {
		tx.Rollback()
		log.Err(err)
		return err
	}

	statement := "INSERT INTO TRANSACTION VALUES($1, $2, $3, $4, $5, $6)"
	if _, err := tx.ExecContext(
		ctx,
		statement,
		t.ID,
		t.AccountID,
		t.OperationType,
		t.Balance,
		t.Amount,
		t.EventDate,
	); err != nil {
		log.Err(err).Msg("cannot add transaction into database")
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Err(err)
		return err
	}

	return nil
}
