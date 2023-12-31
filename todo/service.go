package todo

import (
	"context"
	"fairusatoir/simple-to-do/todo/domain"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func ListAll(ctx context.Context) []domain.Task {

	tx, err := SetPool().Begin()
	PanicIfError(err)
	defer CommitOrRollback(tx)

	return All(ctx, *tx)
}

func SaveItem(ctx context.Context, item domain.Task) domain.Task {
	err := validator.New().Struct(item)
	PanicIfError(err)

	tx, err := SetPool().Begin()
	PanicIfError(err)
	defer CommitOrRollback(tx)

	return Insert(ctx, *tx, item)
}
