package services

import (
	"context"
	"database/sql"
	"simple-to-do/internal/model"
	"simple-to-do/internal/repositories"
	"simple-to-do/internal/utils"

	"github.com/go-playground/validator/v10"
)

type TodoService struct {
	Repo  repositories.Repositories
	dm    *sql.DB
	Valid *validator.Validate
}

func InitalizeTodoService(r repositories.Repositories, db *sql.DB, v *validator.Validate) Service {
	return &TodoService{
		Repo:  r,
		dm:    db,
		Valid: v,
	}
}

func (ts *TodoService) FindAll(c context.Context) ([]model.Task, error) {
	tx, err := ts.dm.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	t, err := ts.Repo.All(c, tx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (ts *TodoService) FindByID(c context.Context, id int) (model.Task, error) {
	tx, err := ts.dm.Begin()
	var _t model.Task
	if err != nil {
		return _t, err
	}
	defer utils.CommitOrRollback(tx)

	t, err := ts.Repo.Find(c, tx, id)
	if err == nil {
		return t, nil
	}
	return _t, err
}

func (ts *TodoService) Create(c context.Context, t model.Task) (model.Task, error) {
	if err := ts.Valid.Struct(t); err != nil {
		return model.Task{}, err
	}

	tx, err := ts.dm.Begin()
	if err != nil {
		return model.Task{}, err
	}
	defer utils.CommitOrRollback(tx)

	nt, err := ts.Repo.Save(c, tx, t)
	if err != nil {
		return model.Task{}, err
	}

	return nt, nil
}

func (ts *TodoService) Update(c context.Context, t model.Task) (model.Task, error) {
	err := ts.Valid.Struct(t)
	var _t model.Task
	if err != nil {
		return _t, err
	}

	tx, err := ts.dm.Begin()
	if err != nil {
		return _t, err
	}
	defer utils.CommitOrRollback(tx)

	_, err = ts.Repo.Find(c, tx, t.Id)
	if err != nil {
		return _t, err
	}

	nt, err := ts.Repo.Update(c, tx, t)
	if err != nil {
		return _t, err
	}

	return nt, nil
}

func (ts *TodoService) Delete(c context.Context, id int) error {
	tx, err := ts.dm.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	err = ts.Repo.Delete(c, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TodoService) UpdateStatus(c context.Context, id int, s bool) (model.Task, error) {
	var _t model.Task

	tx, err := ts.dm.Begin()
	if err != nil {
		return _t, err
	}
	defer utils.CommitOrRollback(tx)

	t, err := ts.Repo.Find(c, tx, id)
	if err != nil {
		return _t, err
	}

	if t.UpdateCompleted(s) {
		t, err = ts.Repo.Update(c, tx, t)
		if err != nil {
			return _t, err
		}
	}

	return t, nil
}
