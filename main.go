package main

import (
	"task-cli/service"
	"task-cli/cli"
)

func main() {
	taskService := service.TaskService{}

	cli.Cli(&taskService)
}
