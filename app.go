package main

import (
	"log"
	"os"
	"os/exec"

	fiber "github.com/gofiber/fiber/v2"
)

const port = ":3000"

func main() {

	// clear cmd
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	// Instantiate fiber app
	app := fiber.New()

	// Routeing
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/users", getUsers)

	app.Use(notFound)

	// Create server
	log.Fatal(app.Listen(port))
}

func getUsers(c *fiber.Ctx) error {
	return c.SendString("All users!")
}

func notFound(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
