package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	router "github.com/smitpatelx/neurelo-go-htmx-example/src/router"
)

func main() {
	// Env File
	env := os.Getenv("ENV")
	if env == "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Print("Could not load .env file")
		}
	} else {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Print("Could not load .env.local file")
		}
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ALLOWED"),
		AllowCredentials: true,
		AllowMethods:     "GET, PUT, POST, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}))

	// Register all routes
	router.SetupRoutes(app)

	// Start server
	serverUrl := os.Getenv("SERVER_URL")
	app.Listen(serverUrl)
}
