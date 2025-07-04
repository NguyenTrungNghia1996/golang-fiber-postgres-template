package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang-fiber-postgres-template/config"
	"golang-fiber-postgres-template/models"
)

func main() {
	app := fiber.New()

	// Kết nối database
	config.ConnectDatabase()

	// Auto migrate bảng User
	models.AutoMigrate(config.DB)

	// Route kiểm tra
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API chạy ngon với module golang-fiber-postgres-template!"})
	})

	// Route tạo user
	app.Post("/users", func(c *fiber.Ctx) error {
		var u models.User
		if err := c.BodyParser(&u); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := config.DB.Create(&u).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(u)
	})

	log.Fatal(app.Listen(":3000"))
}
