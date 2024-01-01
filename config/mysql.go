package config

import (
	"database/sql"
	"simple-to-do/utilities"
	"time"
	
)

func InitMysqlMasterData() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/datamaster?parseTime=true")
	utilities.PanicOnError(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
