package controllers

import (
	"github.com/dduafa/go-server/core/utils"
	"github.com/dduafa/go-server/models"
	"github.com/dduafa/go-server/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userController struct {
	services services.Services
}

func newUserController(s services.Services) *userController {
	return &userController{
		services: s,
	}
}

func (u *userController) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Invalid request body", Data: &fiber.Map{"data": err.Error()}})
	}
	if err := u.services.Users.CreateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(fiber.StatusOK).JSON(utils.CommonResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func (u *userController) FindAllUsers(c *fiber.Ctx) error {
	users, err := u.services.Users.FindAllUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(fiber.StatusOK).JSON(utils.CommonResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": users}})
}

func (u *userController) FindUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Invalid ID", Data: &fiber.Map{"data": err.Error()}})
	}
	user, err := u.services.Users.FindUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(fiber.StatusOK).JSON(utils.CommonResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func (u *userController) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Invalid ID", Data: &fiber.Map{"data": err.Error()}})
	}
	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Invalid request body", Data: &fiber.Map{"data": err.Error()}})
	}
	user, err := u.services.Users.FindUserByID(id)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusNotFound, Message: "User does not exist", Data: &fiber.Map{"data": err.Error()}})
	}
	if err := u.services.Users.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Failed updating user", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(fiber.StatusOK).JSON(utils.CommonResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

func (u *userController) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "Invalid ID", Data: &fiber.Map{"data": err.Error()}})
	}
	if err := u.services.Users.DeleteUserById(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(fiber.StatusOK).JSON(utils.CommonResponse{Status: fiber.StatusOK, Message: "User deleted", Data: &fiber.Map{"data": nil}})
}
