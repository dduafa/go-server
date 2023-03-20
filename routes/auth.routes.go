package routes

import (
	"github.com/dduafa/go-server/controllers"
	"github.com/dduafa/go-server/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, services services.Services) {
	c := controllers.NewController(services)

	app.Post("/register", c.Auth.UserSignUp)
}
