package main

import (
	"task-cli/cli"
	"task-cli/service"
	"task-cli/repository"
)

func main() {
	taskRepo := &repository.InMemoryTaskRepository{}
	taskService := service.NewTaskService(taskRepo)
	cli.Cli(taskService)
}
