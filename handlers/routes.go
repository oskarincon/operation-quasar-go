package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oskarincon/operation-quasar-go/constants"
	"github.com/oskarincon/operation-quasar-go/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post(constants.Apipath+"/topsecret", controllers.PostTopSecret)
	app.Post(constants.Apipath+"/topsecret_split/:satellite_name", controllers.PostTopSecretSplit)
	app.Get(constants.Apipath+"/topsecret_split", controllers.GetTopSecretSplit)
}
