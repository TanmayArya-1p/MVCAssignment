package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	err := app.Listen(conf.Config.INORDER_PORT)
	if err != nil {
		log.Fatal(err)
	}

}
