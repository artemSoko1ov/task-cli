package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func addTask(title string, id int, tasks []Task) []Task {
	task := Task{
		ID:        id,
		Title:     title,
		Completed: false,
	}

	return append(tasks, task)
}

func listTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "    ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func completeTask(id int, tasks []Task) []Task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
		}
	}

	return tasks
}

func deleteTask(id int, tasks []Task) []Task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	return tasks
}

func main() {
	fmt.Println("=====|TASK TRACKER|=====")

	tasks := []Task{}

	scanner := bufio.NewScanner(os.Stdin)

	nextID := 0

	for scanner.Scan() {
		input := scanner.Text()

		command := strings.Split(input, " ")

		switch command[0] {
		case "add":
			if len(command) < 2 {
				fmt.Println("Укажите название задачи")
				return
			}

			title := strings.Join(command[1:], " ")
			fmt.Println("Добавляем задачу: ", title)
			nextID++
			tasks = addTask(title, nextID, tasks)

		case "list":
			listTasks(tasks)

		case "complete":
			if len(command) < 2 {
				fmt.Println("Укажите номер задачи")
				return
			}

			id := command[1]

			number, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Номер задачи должен быть числом")
				return
			}

			fmt.Printf("Завершаем %v задачу\n", number)
			tasks = completeTask(number, tasks)
		
		case "delete":
			if len(command) < 2 {
				fmt.Println("Укажите номер задачи")
				return
			}

			id := command[1]

			number, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Номер задачи должен быть числом")
				return
			}

			fmt.Printf("Удаляем %v задачу\n", number)
			tasks = deleteTask(number, tasks)
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
