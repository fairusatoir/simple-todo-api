package repositories

import (
	"context"
	"database/sql"
	"simple-to-do/internal/model"
	"simple-to-do/internal/utils/constants"
)

type TodoDatamaster struct {
}

func InitalizeTodoDatamaster() Repositories {
	return &TodoDatamaster{}
}

func (td *TodoDatamaster) All(c context.Context, tx *sql.Tx) ([]model.Task, error) {
	q := "SELECT id, title, description, due_date, is_completed FROM task"
	rs, err := tx.QueryContext(c, q)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var ts []model.Task
	for rs.Next() {
		var t model.Task
		if err := rs.Scan(&t.Id, &t.Title, &t.Description, &t.DueDate, &t.IsCompleted); err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	if err = rs.Err(); err != nil {
		return nil, err
	}

	return ts, nil
}

func (td *TodoDatamaster) Find(c context.Context, tx *sql.Tx, id int) (model.Task, error) {
	var t model.Task

	q := "SELECT id, title, description, due_date, is_completed FROM task WHERE id = ?"
	r, err := tx.QueryContext(c, q, id)
	if err != nil {
		return t, err
	}
	defer r.Close()

	if r.Next() {
		err := r.Scan(&t.Id, &t.Title, &t.Description, &t.DueDate, &t.IsCompleted)
		if err != nil {
			return t, err
		}
		return t, nil
	} else {
		return t, constants.Err404
	}
}

func (td *TodoDatamaster) Save(c context.Context, tx *sql.Tx, t model.Task) (model.Task, error) {
	var _t model.Task

	q := "INSERT INTO task(title, description, due_date) VALUES (?,?,?)"
	r, err := tx.ExecContext(c, q, t.Title, t.Description, t.DueDate)
	if err != nil {
		return _t, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return _t, err
	}

	t.Id = int(id)

	return t, nil
}

func (td *TodoDatamaster) Update(c context.Context, tx *sql.Tx, t model.Task) (model.Task, error) {
	q := "UPDATE task SET title = ?, description = ?, due_date = ?, is_completed = ? WHERE id = ?"
	r, err := tx.ExecContext(c, q, t.Title, t.Description, t.DueDate, t.IsCompleted, t.Id)

	var _t model.Task

	if err != nil {
		return _t, err
	}

	_, err = r.RowsAffected()
	if err != nil {
		return _t, err
	}

	return t, nil
}

func (td *TodoDatamaster) Delete(c context.Context, tx *sql.Tx, id int) error {
	q := "DELETE FROM task WHERE id = ?"
	r, err := tx.ExecContext(c, q, id)

	if err != nil {
		return err
	}

	_, err = r.RowsAffected()
	if err == nil {
		return err
	}

	return nil
}
