package main

import (
	"fmt"
	"log"

	"github.com/dduafa/go-server/core"
	"github.com/gofiber/fiber/v2"

	"github.com/dduafa/go-server/core/database"
	"github.com/dduafa/go-server/repositories"
	"github.com/dduafa/go-server/routes"
	"github.com/dduafa/go-server/services"
)

func main() {
	app := fiber.New()
	config := core.NewConfig()

	dbInstance, err := database.NewDBInstance(config)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
	}

	repo := repositories.NewRepository(dbInstance)
	service := services.NewService(repo, config)
	router := routes.NewRouter(app, config, service)

	router.RegisterRoutes()

	log.Fatal(app.Listen(":" + fmt.Sprintf("%v", config.SERVER_PORT)))
}
