package repository

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
	"transaction-manager/database"
)

type Account interface {
	Save(a AccountEntity) error
	FindOne(ID uuid.UUID) (AccountEntity, error)
}

type AccountRepo struct {
	db database.DBConnection
}

type AccountEntity struct {
	ID             uuid.UUID
	DocumentNumber string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewAccount(db database.DBConnection) AccountRepo {
	return AccountRepo{db: db}
}

func (ar AccountRepo) Save(a AccountEntity) error {
	db := ar.db.Connect()
	defer db.Close()

	statement := "INSERT INTO account VALUES($1, $2, $3, $4)"
	if _, err := db.Exec(
		statement,
		a.ID,
		a.DocumentNumber,
		a.CreatedAt,
		a.UpdatedAt,
	); err != nil {
		log.Err(err).Msg("cannot save account into database")
		return err
	}

	return nil
}

func (ar AccountRepo) FindOne(ID uuid.UUID) (AccountEntity, error) {
	db := ar.db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM account WHERE id=$1", ID)
	account := AccountEntity{}
	if err := row.Scan(
		&account.ID,
		&account.DocumentNumber,
		&account.CreatedAt,
		&account.UpdatedAt,
	); err != nil {
		log.Err(err).Msg("cannot find one account on database")
		return account, err
	}
	return account, nil
}
