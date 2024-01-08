package utils

import "database/sql"

func CommitOrRollback(tx *sql.Tx) error {
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// func CommitOrRollback(tx *sql.Tx) {
// 	err := recover()
// 	if err != nil {
// 		errRollback := tx.Rollback()
// 		PanicOnError(errRollback)
// 		panic(err)
// 	} else {
// 		errCommit := tx.Commit()
// 		PanicOnError(errCommit)
// 	}
// }
