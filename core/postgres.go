package core

import (
	"time"
	"gorm.io/gorm"
)

type ConnPostgres struct {
	DB *gorm.DB
}

func ConfigureDBPool(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)
}

func GetDBPool(config DatabaseConfig) (*gorm.DB, error) {
	db, err := NewDatabaseConnection(config)
	if err != nil {
		return nil, err
	}

	ConfigureDBPool(db)
	return db, nil
}

func (conn *ConnPostgres) ExecutePreparedQuery(query string, args ...interface{}) error {
	return conn.DB.Exec(query, args...).Error
}

func (conn *ConnPostgres) FetchRows(query string, args ...interface{}) interface{} {
	return conn.DB.Raw(query, args...)
}

