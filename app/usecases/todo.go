package usecases

import (
	"context"
	"simple-to-do/app/domains"
	"simple-to-do/utilities"
)

func (u *usecase) GetItems(c context.Context) ([]domains.Task, error) {
	tx, err := u.MasterdataClient.Begin()
	if err != nil {
		return nil, err
	}
	defer utilities.CommitOrRollback(tx)

	tasks, err := u.Repo.All(c, tx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u *usecase) GetItemById(c context.Context, id int) (domains.Task, error) {
	tx, err := u.MasterdataClient.Begin()
	if err != nil {
		return domains.Task{}, err
	}
	defer utilities.CommitOrRollback(tx)

	task, err := u.Repo.Find(c, tx, id)
	if err != nil {
		return domains.Task{}, err
	}

	if task.Id == 0 {
		panic(utilities.NewNotFoundError(id))
	}

	return task, nil
}

func (u *usecase) InsertItem(c context.Context, item domains.Task) (domains.Task, error) {
	err := u.Validate.Struct(item)
	if err != nil {
		return domains.Task{}, err
	}

	tx, err := u.MasterdataClient.Begin()
	if err != nil {
		return domains.Task{}, err
	}
	defer utilities.CommitOrRollback(tx)

	NewItem, err := u.Repo.Save(c, tx, item)

	if err != nil {
		return domains.Task{}, err
	}

	return NewItem, nil
}

func (u *usecase) UpdateItem(c context.Context, item domains.Task) (domains.Task, error) {
	err := u.Validate.Struct(item)
	if err != nil {
		return domains.Task{}, err
	}

	tx, err := u.MasterdataClient.Begin()
	if err != nil {
		return domains.Task{}, err
	}
	defer utilities.CommitOrRollback(tx)

	task, err := u.Repo.Find(c, tx, item.Id)
	if err != nil {
		return domains.Task{}, err
	}

	if task.Id == 0 {
		panic(utilities.NewNotFoundError(item.Id))
	}

	NewItem, err := u.Repo.Update(c, tx, item)
	if err != nil {
		return domains.Task{}, err
	}

	return NewItem, nil
}

func (u *usecase) DeleteItem(c context.Context, id int) error {
	tx, err := u.MasterdataClient.Begin()
	if err != nil {
		return err
	}
	defer utilities.CommitOrRollback(tx)

	err = u.Repo.Delete(c, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) UpdateCompletedItem(c context.Context, item domains.UpdateStatusTask) (domains.Task, error) {
	err := u.Validate.Struct(item)
	if err != nil {
		return domains.Task{}, err
	}

	tx, err := u.MasterdataClient.Begin()
	if err != nil {
		return domains.Task{}, err
	}
	defer utilities.CommitOrRollback(tx)

	task, err := u.Repo.Find(c, tx, item.Id)
	if err != nil {
		return domains.Task{}, err
	}

	if task.Id == 0 {
		panic(utilities.NewNotFoundError(item.Id))
	}

	if task.IsCompleted != item.IsCompleted {
		task.IsCompleted = item.IsCompleted
		task, err = u.Repo.Update(c, tx, task)
		if err != nil {
			return domains.Task{}, err
		}
	}

	return task, nil
}
