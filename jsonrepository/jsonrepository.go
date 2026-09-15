package jsonrepository

import (
	"encoding/json"
	"errors"
	"os"
	"task-cli/task"
)

type JSONTaskRepository struct {
	tasks    []task.Task
	filePath string
}

func NewJSONTaskRepository(filePath string) (*JSONTaskRepository, error) {
	repo := &JSONTaskRepository{
		tasks:    []task.Task{},
		filePath: filePath,
	}

	err := repo.load()
	if err != nil {
		return nil, err
	}

	return repo, nil
}


func (j *JSONTaskRepository) load() error {
	_, err := os.Stat(j.filePath)

	if err != nil {
		if os.IsNotExist(err) {
			j.tasks = []task.Task{}
			return j.save()
		}

		return errors.New("Ошибка проверки файла")
	}

	data, err := os.ReadFile(j.filePath)
	if err != nil {
		return errors.New("Ошибка чтения файла")
	}

	err = json.Unmarshal(data, &j.tasks)
	if err != nil {
		return errors.New("Ошибка чтения JSON")
	}

	return nil
}

func (j *JSONTaskRepository) save() error {
	data, err := json.MarshalIndent(j.tasks, "", "    ")
	if err != nil {
		return errors.New("Ошибка создания JSON")
	}

	err = os.WriteFile(j.filePath, data, 0644)
	if err != nil {
		return errors.New("Ошибка записи файла")
	}

	return nil
}

func (j *JSONTaskRepository) Add(t task.Task) error {
	j.tasks = append(j.tasks, t)
	return j.save()
}

func (j *JSONTaskRepository) GetAll() []task.Task {
	return j.tasks
}

func (j *JSONTaskRepository) Complete(id int) error {
	for i := range j.tasks {
		if j.tasks[i].ID == id {
			j.tasks[i].Completed = true
			return j.save()
		}
	}

	return errors.New("Задачи с таким ID не существует")
}

func (j *JSONTaskRepository) Delete(id int) error {
	for i := range j.tasks {
		if j.tasks[i].ID == id {
			j.tasks = append(j.tasks[:i], j.tasks[i+1:]...)
			return j.save()
		}
	}

	return errors.New("Задачи с таким ID не существует")
}