# Todo Backend API (GoFiber)

This project is a Todo backend API built using **Go** and the **Fiber framework**.

The API allows users to create, delete, complete, and retrieve todos with persistence using an in-memory store backed by a JSON file.

---

## Features

- Create a todo
- Delete a todo
- Mark todo as completed
- Fetch all todos
- Live statistics (total, completed, pending)
- Persistence using JSON file
- Constraint: only one todo can be completed at a time
- Thread-safe updates using mutex

---

## Tech Stack

- Go
- Fiber (HTTP framework)
- UUID for unique IDs
- JSON file persistence



---

## Run Locally

### Install dependencies

go get github.com/gofiber/fiber/v2
go get github.com/google/uuid

---

### Start server

go run main.go
