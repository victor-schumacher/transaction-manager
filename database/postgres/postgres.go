package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"transaction-manager/config"

	_ "github.com/lib/pq"
)

const dataSourcePattern = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

type Manager struct {
	c config.Config
}

func NewConnection(c config.Config) Manager {
	return Manager{c: c}
}

// Connect opens postgres connection, using uri provided on config, panics if error
func (m Manager) Connect() *sql.DB {
	log.Println(m.dataSourceName())

	db, err := sql.Open("postgres", m.dataSourceName())
	if err != nil {
		log.Panic(err)
	}
	return db
}

func (m Manager) dataSourceName() string {
	return fmt.Sprintf(
		dataSourcePattern,
		m.c.Database.Host,
		m.c.Database.Port,
		m.c.Database.Username,
		m.c.Database.Password,
		m.c.Database.Name,
		m.c.Database.SSLMode,
	)
}

// TestConnection pings the database, used to ensure database connection is ok, panics if error
func (m Manager) TestConnection() {
	c := m.Connect()
	if err := c.Ping(); err != nil {
		log.Panic(err)
	}
}
