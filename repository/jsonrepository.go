package jsonrepository

import (
	"os"
	"encoding/json"
	"errors"
	"task-cli/task"
)

type JSONTaskRepository struct {
    tasks    []task.Task
    filePath string
}

func (j *JSONTaskRepository) load() error {
	data, err := os.ReadFile(j.filePath)
 	
	if err != nil {
		return errors.New("Ошибка чтения файла")
	}

	 err = json.Unmarshal(data, &j.tasks)

	 if err != nil {
		return errors.New("Ошибка ")
	}

	return nil
}

func (j *JSONTaskRepository) save() error {
	data, err := json.MarshalIndent(j.tasks, "", "    ")

	if err != nil {
		return errors.New("Ошибка ")
	}

	err = os.WriteFile(j.filePath, data)
 	
	if err != nil {
		return errors.New("Ошибка записи файла")
	}

	return nil
}