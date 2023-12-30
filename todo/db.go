package todo

import (
	"database/sql"
	"time"
)

// https://github.com/go-sql-driver/mysql/

func SetPool() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/datamaster?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
