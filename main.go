package main

import (
	"creztfallen/go-rest-api/configs"
	"creztfallen/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	
	configs.ConnectDB()

	routes.UserRoute((app))

	app.Listen(":6000")
}