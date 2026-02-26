package storage

import (
	"encoding/json"
	"os"
	"sync"
	"todo-test/models"
)

var (
	Todos []models.Todo
	Mutex sync.Mutex
	File  = "data/todos.json"
)

func Load() error {
	data, err := os.ReadFile(File)
	if err != nil {
		return nil
	}
	return json.Unmarshal(data, &Todos)
}

func Save() error {
	data, err := json.MarshalIndent(Todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(File, data, 0644)
}