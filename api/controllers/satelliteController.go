package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oskarincon/operation-quasar-go/models"
	"github.com/oskarincon/operation-quasar-go/services"
)

func PostTopSecret(c *fiber.Ctx) error {
	data := new(models.Satellites)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	fmt.Printf("[PostTopSecret] data: %#v\n", data.Satellites)
	res, err := services.FindPositionMessage(*data)
	fmt.Printf("[PostTopSecret] - FindPositionMessage res: %#v\n, pos: %#v\n", res, err)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.JSON(res)
}

func PostTopSecretSplit(c *fiber.Ctx) error {
	data := new(models.Satellite)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	satelliteName := c.Params("satellite_name")
	fmt.Printf("[PostTopSecretSplit] satelliteName: %#v\n", satelliteName)
	res, err := services.PostInfoSatellite(satelliteName, *data)
	fmt.Printf("[PostTopSecretSplit] - PostInfoSatellite res: %#v\n, pos: %#v\n", res, err)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.JSON(res)
}

func GetTopSecretSplit(c *fiber.Ctx) error {
	dataSatellite, err := services.GetInfoSatellite()
	fmt.Printf("[GetTopSecretSplit] - GetInfoSatellite res: %#v\n, pos: %#v\n", dataSatellite, err)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	res, err := services.FindPositionMessage(dataSatellite)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "data": err.Error()})
	}
	return c.JSON(res)
}
