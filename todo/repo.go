package todo

import (
	"context"
	"database/sql"
	"errors"
	"fairusatoir/simple-to-do/todo/domain"
	"net/http"
)

func All(ctx context.Context, tx sql.Tx) []domain.Task {

	q := "SELECT id, title, description, due_date, is_completed FROM task"
	rows, err := tx.QueryContext(ctx, q)
	PanicIfError(err)
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		task := domain.Task{}
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.IsCompleted)
		PanicIfError(err)
		tasks = append(tasks, task)
	}

	return tasks
}

func Find(ctx context.Context, tx sql.Tx, id int) (domain.Task, error) {

	q := "SELECT id, title, description, due_date, is_completed FROM task WHERE id=?"
	rows, err := tx.QueryContext(ctx, q, id)
	PanicIfError(err)
	defer rows.Close()

	task := domain.Task{}
	if rows.Next() {
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.IsCompleted)
		PanicIfError(err)
		return task, nil
	} else {
		return task, errors.New(http.StatusText(404))
	}
}

func Insert(ctx context.Context, tx sql.Tx, t domain.Task) domain.Task {

	q := "INSERT INTO task (title, description, due_date)VALUES(?, ?, ?);"
	r, err := tx.ExecContext(ctx, q, t.Title, t.Description, t.DueDate)
	PanicIfError(err)

	id, err := r.LastInsertId()
	PanicIfError(err)

	t.Id = int(id)
	return t
}
