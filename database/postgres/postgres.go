package postgres

import (
	"database/sql"
	"log"
)

type Manager struct {
}

func NewConnection() Manager {
	return Manager{}
}

// Connect opens postgres connection, using uri provided on config, panics if error
func (m Manager) Connect() *sql.DB {
	db, err := sql.Open("postgres", "")
	if err != nil {
		log.Panic(err)
	}
	return db
}

// TestConnection pings the database, used to ensure database connection is ok, panics if error
func (m Manager) TestConnection() {
	c := m.Connect()
	if err := c.Ping(); err != nil {
		log.Panic(err)
	}
}
