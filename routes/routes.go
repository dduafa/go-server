package routes

import (

	"github.com/dduafa/go-server/core"
	"github.com/dduafa/go-server/responses"
	"github.com/dduafa/go-server/services"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	app      *fiber.App
	config   *core.Config
	services services.Services
}

func NewRouter(app *fiber.App, config *core.Config, services services.Services) *router {
	return &router{
		app:      app,
		config:   config,
		services: services,
	}
}

func (r *router) RegisterRoutes() {
	r.app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(responses.CommonResponse{Status: fiber.StatusOK, Message: "Hi, ✋ Server is Okay!!!"})
	})

	AuthRoutes(r.app, r.services)
}
