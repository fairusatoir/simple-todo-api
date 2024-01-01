package usecases

import (
	"context"
	"database/sql"
	"simple-to-do/app/domains"
	"simple-to-do/app/repositories"

	"github.com/go-playground/validator/v10"
)

type usecase struct {
	Repo             repositories.Repositories
	MasterdataClient *sql.DB
	Validate         *validator.Validate
}

func NewUsecase(r repositories.Repositories, db *sql.DB, v *validator.Validate) *usecase {
	return &usecase{
		MasterdataClient: db,
		Repo:             r,
		Validate:         v,
	}
}

type Usecase interface {
	GetItems(c context.Context) ([]domains.Task, error)
	GetItemById(c context.Context, id int) (domains.Task, error)
	InsertItem(c context.Context, item domains.Task) (domains.Task, error)
	UpdateItem(c context.Context, item domains.Task) (domains.Task, error)
}
