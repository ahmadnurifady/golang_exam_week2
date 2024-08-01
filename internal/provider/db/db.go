package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Connection interface {
	GetDb() *sql.DB
}

type connection struct {
	db *sql.DB
}

func (c *connection) getDb() error {
	driver := "postgres"

	connStr := "user=postgres dbname=golang_database sslmode=disable password=101020"

	DB, err := sql.Open(driver, connStr)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	c.db = DB

	return nil
}

func (c *connection) GetDb() *sql.DB {
	return c.db
}

func NewConnectionDatabase() (Connection, error) {
	conn := &connection{}
	err := conn.getDb()
	if err != nil {
		return nil, err
	}

	return conn, nil

}
