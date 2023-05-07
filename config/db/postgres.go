package db

import (
	"database/sql"
	"fundamental-payroll-go/helper/apperrors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPgxDatabase(driver, url string) (*sql.DB, error) {
	if url == "" {
		return nil, apperrors.New(apperrors.ErrDbUrlNotFound)
	}

	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
