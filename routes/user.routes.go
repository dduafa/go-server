package routes

import (
	"github.com/dduafa/go-server/controllers"
	"github.com/dduafa/go-server/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, services services.Services) {
	c := controllers.NewController(services)

	app.Post("/user", c.User.CreateUser)
	app.Get("/user", c.User.FindAllUsers)
	app.Get("/user/:id", c.User.FindUser)
	app.Put("/user/:id", c.User.UpdateUser)
	app.Delete("/user/:id", c.User.DeleteUser)
}
