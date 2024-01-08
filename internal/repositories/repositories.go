package repositories

import (
	"context"
	"database/sql"
	"simple-to-do/internal/model"
)

type Repositories interface {
	All(c context.Context, tx *sql.Tx) ([]model.Task, error)
	Find(c context.Context, tx *sql.Tx, id int) (model.Task, error)
	Save(c context.Context, tx *sql.Tx, t model.Task) (model.Task, error)
	Update(c context.Context, tx *sql.Tx, t model.Task) (model.Task, error)
	Delete(c context.Context, tx *sql.Tx, id int) error
}
