package main

import (
	"os"
	"os/exec"

	"github.com/gofiber/fiber"
)

func main() {

	// clear cmd
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// Instantiate fiber app
	app := fiber.New()

	// Routeing
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello")
	})

	app.Use(notFound)

	// Create server
	app.Listen(3000)
}

func notFound(c *fiber.Ctx) {
	c.SendStatus(404)
}
