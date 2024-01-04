package usecases

import (
	"context"
	"simple-to-do/app/domains"
)

type Usecase interface {
	GetItems(c context.Context) ([]domains.Task, error)
	GetItemById(c context.Context, id int) (domains.Task, error)
	InsertItem(c context.Context, item domains.Task) (domains.Task, error)
	UpdateItem(c context.Context, item domains.Task) (domains.Task, error)
	DeleteItem(c context.Context, id int) error
	UpdateCompletedItem(c context.Context, item domains.UpdateStatusTask) (domains.Task, error)
}
