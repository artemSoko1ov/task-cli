package service

import (
	"task-cli/task"
	"task-cli/repository"
)

type TaskService struct {
	repo  repository.TaskRepository
	nextID int
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
   return  &TaskService{
    	repo: repo,
	}
}

func (s *TaskService) AddTask(title string) {
	task := task.Task{
		ID:        s.incrementId(),
		Title:     title,
		Completed: false,
	}

	s.repo.Add(task)
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
