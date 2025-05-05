package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kartik122/remondata/internal/routes"
)

func main() {
	app := fiber.New()
	routes.Register(app)
	app.Listen(":3000")
}
