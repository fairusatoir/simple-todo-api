package domains

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsCompleted bool      `json:"is_completed" validate:"required_without=Title"`
}

type UpdateStatusTask struct {
	Id          int  `json:"id"`
	IsCompleted bool `json:"is_completed" validate:"required"`
}

func (t *Task) Done() {
	if !t.IsCompleted {
		t.IsCompleted = true
	}
}
