package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-fiber-postgres-template/controllers"
	"golang-fiber-postgres-template/middleware"
	"golang-fiber-postgres-template/models"
)

// Setup registers all application routes.
func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(models.APIResponse{
			Status:  "success",
			Message: "API chạy ngon với module golang-fiber-postgres-template!",
		})
	})

	app.Post("/users", controllers.CreateUser)
	app.Post("/login", controllers.Login)

	api := app.Group("/api", middleware.Protected())
	api.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON(models.APIResponse{
			Status:  "success",
			Message: "authorized",
		})
	})
}
