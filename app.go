package main

import "github.com/gofiber/fiber"

func main() {
	// Fiber instace
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// start server
	app.Listen(3000)
}

func hello(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
