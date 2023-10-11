package main

import (
	"context"
	"log"
	"os"

	"server/db"
	"server/router"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	client, err := db.InitializeMongoDB()

	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	defer client.Disconnect(context.Background())

	port := PORT
	if port == "" {
		port = "3000"
	}

	api := app.Group("/api")

	router.RegisterApiRoutes(api, client)

	log.Printf("Listening on port %s\n", port)

	log.Fatal(app.Listen(":" + port))
}
