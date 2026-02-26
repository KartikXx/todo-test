package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"todo-test/handlers"
	"todo-test/storage"
)

func main() {
	os.MkdirAll("data", os.ModePerm)

	if err := storage.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/todos", handlers.AddTodo)
	app.Get("/todos", handlers.GetTodos)
	app.Delete("/todos/:id", handlers.DeleteTodo)
	app.Put("/todos/:id/complete", handlers.CompleteTodo)
	app.Get("/stats", handlers.Stats)

	log.Fatal(app.Listen(":3000"))
}