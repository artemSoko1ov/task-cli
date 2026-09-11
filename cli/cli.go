package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-cli/service"
)

// func listTasks(tasks []Task) {
// 	data, err := json.MarshalIndent(tasks, "", "    ")

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(data))
// }

func Cli(taskService *service.TaskService) {
	fmt.Println("=====|TASK TRACKER|=====")

	scanner := bufio.NewScanner(os.Stdin)

	nextID := 0

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
			nextID++
			taskService.AddTask(title, nextID)

		case "list":
			// listTasks(tasks)

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

			fmt.Printf("Завершаем %v задачу\n", number)
			taskService.CompleteTask(number)

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

			fmt.Printf("Удаляем %v задачу\n", number)
			taskService.DeleteTask(number)
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
