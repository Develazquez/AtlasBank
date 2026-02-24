package core

import (
	"database/sql"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func CreateDBConnection(config DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error creando conexión: %w", err)
	}

	return db, nil
}


func NewDatabaseConnection(config DatabaseConfig) (*sql.DB, error) {
	db, err := CreateDBConnection(config)
	if err != nil {
		return nil, err
	}


	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verificando conexión: %w", err)
	}

	return db, nil
}
