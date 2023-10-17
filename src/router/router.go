package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// --- API groups   ---
	// api_1 := app.Group("/api/v1")

	// Render homepage
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(os.Getenv("FRONT_END_URL"))
	})
}
