package repositories

import (
	"context"
	"database/sql"
	"simple-to-do/app/domains"
	"simple-to-do/utilities"
)

type Datamaster struct {
}

func NewDatamaster() Repositories {
	return &Datamaster{}
}

func (r *Datamaster) All(c context.Context, tx *sql.Tx) ([]domains.Task, error) {
	q := "SELECT id, title, description, due_date, is_completed FROM task"
	rows, err := tx.QueryContext(c, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domains.Task
	for rows.Next() {
		var item domains.Task
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.DueDate, &item.IsCompleted)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	err = rows.Err() // Periksa error setelah loop selesai
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Datamaster) Find(c context.Context, tx *sql.Tx, id int) (domains.Task, error) {
	q := "SELECT id, title, description, due_date, is_completed FROM task WHERE id = ?"
	row := tx.QueryRowContext(c, q, id)

	var task domains.Task
	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.DueDate, &task.IsCompleted)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(utilities.NewNotFoundError(id))
		}
		return domains.Task{}, err
	}

	return task, nil
}

func (r *Datamaster) Save(c context.Context, tx *sql.Tx, item domains.Task) (domains.Task, error) {
	q := "INSERT INTO task(title, description, due_date) VALUES (?,?,?)"
	result, e := tx.ExecContext(c, q, item.Title, item.Description, item.DueDate)
	if e != nil {
		return domains.Task{}, e
	}

	id, e := result.LastInsertId()
	if e != nil {
		return domains.Task{}, e
	}

	item.Id = int(id)

	return item, nil
}

func (r *Datamaster) Update(c context.Context, tx *sql.Tx, item domains.Task) (domains.Task, error) {
	q := "UPDATE task SET title = ?, description = ?, due_date = ?, is_completed = ? WHERE id = ?"
	result, err := tx.ExecContext(c, q, item.Title, item.Description, item.DueDate, item.IsCompleted, item.Id)
	if err != nil {
		return domains.Task{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domains.Task{}, err
	}

	if rowsAffected == 0 {
		panic(utilities.NewNotFoundError(item.Id))
	}

	return item, nil
}

func (r *Datamaster) Delete(c context.Context, tx *sql.Tx, id int) error {
	q := "DELETE FROM task WHERE id = ?"
	result, err := tx.ExecContext(c, q, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		panic(utilities.NewNotFoundError(id))
	}

	return nil
}
