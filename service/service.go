package service

import "task-cli/task"

type TaskService struct {
    tasks []task.Task
}

func (s *TaskService) AddTask(title string, id int) {
    task := task.Task{
        ID:        id,
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