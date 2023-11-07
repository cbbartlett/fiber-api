package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myproject/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	// Add more routes as needed
}
