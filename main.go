package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang-fiber-postgres-template/config"
	"golang-fiber-postgres-template/models"
	"golang-fiber-postgres-template/routes"
)

func main() {
	app := fiber.New()

	// Kết nối database
	config.ConnectDatabase()

	// Auto migrate bảng User
	models.AutoMigrate(config.DB)

	// Register routes
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
