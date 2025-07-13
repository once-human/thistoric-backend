package routes

import (
	"github.com/gofiber/fiber/v2"
	"thistoric-backend/middleware"
)

func Register(app *fiber.App) {
	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Thistoric 👑",
		})
	})

	// Health check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Protected routes example
	api := app.Group("/api/v1")
	
	// Admin only route
	api.Get("/admin/dashboard", middleware.RequireAuth("admin"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Admin Dashboard",
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	// Vendor and Organizer routes
	api.Get("/vendor/dashboard", middleware.RequireAuth("vendor"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Vendor Dashboard",
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	api.Get("/organizer/dashboard", middleware.RequireAuth("organizer"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Organizer Dashboard",
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	// Multi-role route (accessible by vendors and organizers)
	api.Get("/events", middleware.RequireAuth("vendor", "organizer"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Events List",
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	// TODO: Register actual modules here later
	// e.g. RegisterAuthRoutes(app.Group("/api/v1/auth"))
}
