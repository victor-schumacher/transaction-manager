package repository

import (
	"time"
	"transaction-manager/database"
)

type Account interface {
	Save(a AccountEntity) error
	FindOne(ID string) error
}

type AccountRepo struct {
	db database.DBConnection
}

type AccountEntity struct {
	ID             string
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

	statement := "INSERT INTO ACCOUNT VALUES(?, ?, ?, ?)"
	if _, err := db.Exec(
		statement,
		a.ID,
		a.DocumentNumber,
		a.CreatedAt,
		a.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (ar AccountRepo) FindOne(ID string) (AccountEntity, error) {
	db := ar.db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM ACCOUNT WHERE ID=?", ID)
	account := AccountEntity{}
	if err := row.Scan(&account); err != nil {
		return account, err
	}

	return account, nil
}
