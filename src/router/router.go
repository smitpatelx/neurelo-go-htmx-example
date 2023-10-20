package router

import (
	"github.com/gofiber/fiber/v2"
	actor "github.com/smitpatelx/neurelo-go-htmx-example/src/services/actor"
)

func SetupRoutes(app *fiber.App) {
	// --- Render pages ---
	app.Get("/", actor.RenderIndexPage)
	app.Get("/actors", actor.GetAllActors)
	app.Get("/films", actor.GetAllFilms)

	// --- API groups   ---
	api_1 := app.Group("/api/v1")

	// --- Actors API ---
	api_1.Get("/actors", actor.GetAllActorsAPI)
}
