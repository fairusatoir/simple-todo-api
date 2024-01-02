package repositories

import (
	"context"
	"database/sql"
	"simple-to-do/app/domains"
)

type repositories struct {
}

func NewRepositories() *repositories {
	return &repositories{}
}

type Repositories interface {
	All(c context.Context, tx *sql.Tx) ([]domains.Task, error)
	Find(c context.Context, tx *sql.Tx, id int) (domains.Task, error)
	Save(c context.Context, tx *sql.Tx, item domains.Task) (domains.Task, error)
	Update(c context.Context, tx *sql.Tx, item domains.Task) (domains.Task, error)
	Delete(c context.Context, tx *sql.Tx, id int) error
}
