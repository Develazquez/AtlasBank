package core

import (
	"database/sql"
	"time"
)

type ConnPostgres struct {
	DB *sql.DB
}

func ConfigureDBPool(db *sql.DB) {
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)
}

func GetDBPool(config DatabaseConfig) (*sql.DB, error) {
	db, err := NewDatabaseConnection(config)
	if err != nil {
		return nil, err
	}

	ConfigureDBPool(db)
	return db, nil
}

func (conn *ConnPostgres) ExecutePreparedQuery(query string, args ...interface{}) error {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	return err
}

func (conn *ConnPostgres) FetchRows(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
