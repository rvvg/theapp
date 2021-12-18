package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Connect with database
	if err := ConnectDB(); err != nil {
		log.Fatal(err)
	}

	// check if query looks like 'number quals x' and responds to the URL like 'http://host/?n=x' and return n*n
	app.Get("/", func(c *fiber.Ctx) error {
		n := c.Query("n")
		if num, err := strconv.Atoi(n); err == nil {
			log.Default().Printf("%s is a number", n)
			return c.Status(200).SendString(strconv.Itoa(num * num))
		}
		log.Default().Printf("%s is not a number", n)
		return c.Status(422).SendString("Invalid query" + c.IP())
	})

	app.Get("/blacklisted", func(c *fiber.Ctx) error {
		BlacklistIP(c.IP(), c.Path())
		log.Default().Println("IP: " + c.IP() + " blacklisted")
		return c.Status(444).SendString("Blacklisted")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		log.Default().Println("Health check")
		return c.Status(200).SendString("OK")
	})

	app.Listen(":3000")
}
