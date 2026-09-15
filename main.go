package main

import (
	"fmt"
	"task-cli/cli"
	"task-cli/jsonrepository"
	"task-cli/service"
)

func main() {
	taskRepo, err := jsonrepository.NewJSONTaskRepository("tasks.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	taskService := service.NewTaskService(taskRepo)

	cli.Cli(taskService)
}
