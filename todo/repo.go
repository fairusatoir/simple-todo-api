package todo

import (
	"context"
	"database/sql"
	"fairusatoir/simple-to-do/todo/domain"
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

func Insert(ctx context.Context, tx sql.Tx, t domain.Task) domain.Task {

	q := "INSERT INTO task (title, description, due_date)VALUES(?, ?, ?);"
	r, err := tx.ExecContext(ctx, q, t.Title, t.Description, t.DueDate)
	PanicIfError(err)

	id, err := r.LastInsertId()
	PanicIfError(err)

	t.Id = int(id)
	return t
}
