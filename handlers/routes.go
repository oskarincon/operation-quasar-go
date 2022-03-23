package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oskarincon/operation-quasar-go/controllers"
)

func SetupRoutes(router fiber.Router) {
	router.Post("/topsecret", controllers.PostTopSecret)
	router.Post("/topsecret_split/:satellite_name", controllers.PostTopSecretSplit)
	router.Get("/topsecret_split", controllers.GetTopSecretSplit)
}
