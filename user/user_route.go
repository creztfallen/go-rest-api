package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", CreateUser)
	app.Get("/user/:userId", GetAUser)
	app.Put("/user/:userId", EditAUser)
	app.Delete("/user/:userId", DeleteAUser)
	app.Get("/users", GetAllUsers)
}
