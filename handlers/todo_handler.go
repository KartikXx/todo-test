package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
	"todo-test/models"
	"todo-test/storage"
)

func AddTodo(c *fiber.Ctx) error {
	var body struct {
		Title string `json:"title"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	todo := models.Todo{
		ID:        uuid.NewString(),
		Title:     body.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	storage.Mutex.Lock()
	storage.Todos = append(storage.Todos, todo)
	storage.Save()
	storage.Mutex.Unlock()

	return c.JSON(todo)
}

func GetTodos(c *fiber.Ctx) error {
	return c.JSON(storage.Todos)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	for i, t := range storage.Todos {
		if t.ID == id {
			storage.Todos = append(storage.Todos[:i], storage.Todos[i+1:]...)
			storage.Save()
			return c.JSON(fiber.Map{"message": "deleted"})
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "not found"})
}

func CompleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	for i := range storage.Todos {
		if storage.Todos[i].ID == id {

			// only one completed rule
			for j := range storage.Todos {
				storage.Todos[j].Completed = false
			}

			storage.Todos[i].Completed = true
			storage.Save()
			return c.JSON(storage.Todos[i])
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "not found"})
}

func Stats(c *fiber.Ctx) error {
	total := len(storage.Todos)
	completed := 0

	for _, t := range storage.Todos {
		if t.Completed {
			completed++
		}
	}

	return c.JSON(fiber.Map{
		"total":     total,
		"completed": completed,
		"pending":   total - completed,
	})
}