package main

import (
	"go_base_rest_api/app"
	"go_base_rest_api/configs"
	"go_base_rest_api/repository"
	"go_base_rest_api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")
	TodoRepositoryDb := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Get("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":9000")
}
