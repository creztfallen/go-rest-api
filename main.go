package main

import (
	"creztfallen/go-rest-api/configs"
	"creztfallen/go-rest-api/user"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	user.UserRoute((app))

	app.Listen(":6000")
}
