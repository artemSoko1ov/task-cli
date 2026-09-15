package service

import (
	"task-cli/repository"
	"task-cli/task"
)

type TaskService struct {
	repo   repository.TaskRepository
	nextID int
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) AddTask(title string) error {
	t := task.Task{
		ID:        s.incrementId(),
		Title:     title,
		Completed: false,
	}

	return s.repo.Add(t)
}

func (s *TaskService) CompleteTask(id int) error {
	return s.repo.Complete(id)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}

func (s *TaskService) ShowTasks() []task.Task {
	return s.repo.GetAll()
}

func (s *TaskService) incrementId() int {
	s.nextID++
	return s.nextID
}
