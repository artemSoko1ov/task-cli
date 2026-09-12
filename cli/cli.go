package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-cli/service"
	"task-cli/task"
)

func listTasks(tasks []task.Task) {
	data, err := json.MarshalIndent(tasks, "", "    ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func Cli(taskService *service.TaskService) {
	fmt.Println("=====|TASK TRACKER|=====")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		command := strings.Split(input, " ")

		switch command[0] {
		case "add":
			if len(command) < 2 {
				fmt.Println("Укажите название задачи")
				continue
			}

			title := strings.Join(command[1:], " ")
			fmt.Println("Добавляем задачу: ", title)

			taskService.AddTask(title)

		case "list":
			tasks := taskService.ShowTasks()
			listTasks(tasks)

		case "complete":
			if len(command) < 2 {
				fmt.Println("Укажите номер задачи")
				continue
			}

			id := command[1]

			number, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Номер задачи должен быть числом")
				continue
			}

			err = taskService.CompleteTask(number)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Завершаем %v задачу\n", number)
		case "delete":
			if len(command) < 2 {
				fmt.Println("Укажите номер задачи")
				continue
			}

			id := command[1]

			number, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Номер задачи должен быть числом")
				continue
			}

			err = taskService.DeleteTask(number)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Удаляем %v задачу\n", number)

		case "exit":
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
