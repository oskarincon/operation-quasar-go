package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init(app *fiber.App) {
	// Middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Start server on port 3333
	app.Listen(":3333")
}
