package services

import (
	"context"
	"simple-to-do/internal/model"
)

type Service interface {
	FindAll(c context.Context) ([]model.Task, error)
	FindByID(c context.Context, id int) (model.Task, error)
	Create(c context.Context, t model.Task) (model.Task, error)
	Update(c context.Context, t model.Task) (model.Task, error)
	Delete(c context.Context, id int) error
	UpdateStatus(c context.Context, id int, s bool) (model.Task, error)
}
