package repository

import (
	"errors"
	"task-cli/task"
)

type TaskRepository interface {
	Add(t task.Task) error
	GetAll() []task.Task
	Complete(id int) error
	Delete(id int) error
}

type InMemoryTaskRepository struct {
	tasks []task.Task
}

func (r *InMemoryTaskRepository) Add(t task.Task) error {
	r.tasks = append(r.tasks, t)
	return nil
}
func (r *InMemoryTaskRepository) GetAll() []task.Task {
	return r.tasks
}

func (r *InMemoryTaskRepository) Complete(id int) error {
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks[i].Completed = true

			return nil
		}
	}

	return errors.New("Задачи с таким ID не существует")
}

func (r *InMemoryTaskRepository) Delete(id int) error {
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("Задачи с таким ID не существует")
}
