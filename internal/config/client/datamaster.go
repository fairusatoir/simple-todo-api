package client

import (
	"database/sql"
	"fmt"
	"simple-to-do/internal/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitializedDatamaster() (*sql.DB, error) {
	db, err := sql.Open(config.DSDatamasterDriver(), config.DSDatamasterUrl())
	if err != nil {
		return nil, fmt.Errorf("error opening Datamaster: %v", err)
	}
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging Datamaster : %v", err)
	}

	return db, nil
}
