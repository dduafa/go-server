package utils

import "github.com/gofiber/fiber/v2"

type CommonResponse struct {
	Status  int16      `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}