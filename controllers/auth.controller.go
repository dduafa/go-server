package controllers

import (
	"strings"

	"github.com/dduafa/go-server/core/utils"
	"github.com/dduafa/go-server/models"
	"github.com/dduafa/go-server/services"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	services services.Services
}

func newAuthController(s services.Services) *authController {
	return &authController{
		services: s,
	}
}

func (a *authController) UserSignUp(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Invalid request body", Data: &fiber.Map{"data": err.Error()}})
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Password Error", Data: &fiber.Map{"data": err.Error()}})
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
		Photo:    payload.Photo,
	}

	result := a.services.Users.CreateUser(&newUser)

	if result != nil {
		if strings.Contains(result.Error(), "duplicated key not allowed") {
			return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "User with that email already exists", Data: &fiber.Map{"data": err.Error()}})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Error signing up", Data: &fiber.Map{"data": err.Error()}})
		}
	}
	return c.Status(fiber.StatusOK).JSON(utils.CommonResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": newUser}})
}
