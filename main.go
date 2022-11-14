package main

import (
	"log"
	"tobialbertino/be-portfolio/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		cfg *config.Config
		err error
	)

	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + cfg.Server.Port)
}
