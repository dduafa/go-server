package controllers

import (
	"strings"
	"time"

	"github.com/dduafa/go-server/models"
	"github.com/dduafa/go-server/responses"
	"github.com/dduafa/go-server/services"
	"github.com/dduafa/go-server/utils"
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

func (authC *authController) UserSignUp(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}

	user, err := authC.services.Users.FindUserByEmail(payload.Email)

	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{Status: fiber.StatusBadRequest, Message: "User already exists", Data: &fiber.Map{"data": err.Error()}})
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{Status: fiber.StatusBadRequest, Message: "Password Error", Data: &fiber.Map{"data": err.Error()}})
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  true,
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := authC.services.Users.CreateUser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{Status: fiber.StatusBadRequest, Message: "User with that email already exists", Data: &fiber.Map{"data": err.Error()}})
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(fiber.StatusOK).JSON(responses.CommonResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

// func (ac *AuthController) SignInUser(ctx *gin.Context) {
// 	var payload *models.SignInInput

// 	if err := ctx.ShouldBindJSON(&payload); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
// 		return
// 	}

// 	var user models.User
// 	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
// 	if result.Error != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
// 		return
// 	}

// 	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
// 		return
// 	}

// 	config, _ := initializers.LoadConfig(".")

// 	// Generate Tokens
// 	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
// 		return
// 	}

// 	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
// 		return
// 	}

// 	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
// 	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
// 	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
// }

// func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
// 	message := "could not refresh access token"

// 	cookie, err := ctx.Cookie("refresh_token")

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
// 		return
// 	}

// 	config, _ := initializers.LoadConfig(".")

// 	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
// 		return
// 	}

// 	var user models.User
// 	result := ac.DB.First(&user, "id = ?", fmt.Sprint(sub))
// 	if result.Error != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
// 		return
// 	}

// 	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
// 		return
// 	}

// 	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
// 	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
// }

// func (ac *AuthController) LogoutUser(ctx *gin.Context) {
// 	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
// 	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
// 	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

// 	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
// }
