package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kartik122/remondata/internal/handlers"
)

func Register(app *fiber.App) {
	app.Get("/", handlers.Hello)
}
