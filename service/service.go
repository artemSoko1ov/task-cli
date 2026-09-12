package service

import "task-cli/task"

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

func (s *TaskService) CompleteTask(id int) {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Completed = true
		}
	}
}

func (s *TaskService) DeleteTask(id int) {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			break
		}
	}
}

func (s *TaskService) ShowTasks() []task.Task {
	return s.tasks
}

func (s *TaskService) incrementId() int {
	s.nextID++
	return s.nextID
}
