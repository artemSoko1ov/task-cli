package main

import (
	"task-cli/cli"
	"task-cli/service"
)

func main() {
	taskService := service.TaskService{}

	cli.Cli(&taskService)
}
