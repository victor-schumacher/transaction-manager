package database

import "database/sql"

type DBConnection interface {
	Connect() *sql.DB
	TestConnection()
}

type Migration interface {
	Apply()
}
