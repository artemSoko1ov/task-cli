package service

import (
	"errors"
	"task-cli/task"
)

type TaskService struct {
	tasks  []task.Task
	nextID int
}

func (s *TaskService) AddTask(title string) {
	task := task.Task{
		ID:        s.incrementId(),
		Title:     title,
		Completed: false,
	}

	s.tasks = append(s.tasks, task)
}

func (s *TaskService) CompleteTask(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Completed = true

			return nil
		}
	}

	return errors.New("Задачи с таким ID не существует")
}

func (s *TaskService) DeleteTask(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("Задачи с таким ID не существует")
}

func (s *TaskService) ShowTasks() []task.Task {
	return s.tasks
}

func (s *TaskService) incrementId() int {
	s.nextID++
	return s.nextID
}
