package database

import "database/sql"

type DbConnection interface {
	Connect() *sql.DB
	TestConnection()
}

type Migration interface {
	Apply()
}
