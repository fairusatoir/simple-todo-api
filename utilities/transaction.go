package utilities

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicOnError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicOnError(errCommit)
	}
}
