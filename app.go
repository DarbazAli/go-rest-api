package main

import "github.com/gofiber/fiber"

func main() {
	// Fiber instace
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Last middleware that matches anything
	app.Use(notFound)

	// start server
	app.Listen(3000)
}

func hello(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

// not found middlware
func notFound(c *fiber.Ctx) {
	c.SendStatus(404)
}
