package main

import (
	"log"

	"github.com/firebase007/go-rest-api-with-fiber/database"
	"github.com/firebase007/go-rest-api-with-fiber/router"
	"github.com/gofiber/fiber" // import the fiber package
	"github.com/gofiber/fiber/middleware"

	_ "github.com/lib/pq"
)

// Endtry point to begin program execution
func main() {

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// Fiber instace
	app := fiber.New()

	// Middleware
	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	// start server
	app.Listen(3000)
}
