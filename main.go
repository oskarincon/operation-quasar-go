package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oskarincon/operation-quasar-go/handlers"
)

func main() {
	app := fiber.New()

	// Create Routes
	handlers.SetupRoutes(app)
	// Server Init
	handlers.Init(app)
}
