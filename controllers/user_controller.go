package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang-fiber-postgres-template/models"
	"golang-fiber-postgres-template/repositories"
	"golang-fiber-postgres-template/utils"
)

// CreateUser handles user registration.
func CreateUser(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	u.Password = hashed
	if err := repositories.CreateUser(&u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	u.Password = ""
	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{
		Status:  "success",
		Message: "user created",
		Data:    u,
	})
}

// Login authenticates the user and returns a JWT token.
func Login(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	user, err := repositories.GetUserByUsername(req.Username)
	if err != nil || !utils.CheckPassword(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
			Status:  "error",
			Message: "invalid credentials",
		})
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.JSON(models.APIResponse{
		Status:  "success",
		Message: "login successful",
		Data:    fiber.Map{"token": t},
	})
}
