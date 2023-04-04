package db

import (
	"database/sql"
	"errors"
	"fundamental-payroll-go/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPgxDatabase(cfg *config.Config) (*sql.DB, error) {
	if cfg.Database.URL == "" {
		return nil, errors.New("database URL not existed")
	}

	db, err := sql.Open(cfg.Database.Driver, cfg.Database.URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
