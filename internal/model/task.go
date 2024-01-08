package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsCompleted bool      `json:"is_completed" validate:"required_without=Title"`
}

func (t *Task) UpdateCompleted(s bool) bool {
	if t.IsCompleted != s {
		t.IsCompleted = s
		return true
	}
	return false
}

func (t *Task) SetId(id int) {
	t.Id = id
}
