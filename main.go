package main

import (
	"log"
	"os"
	appRouter "tobialbertino/portfolio-be/app"
	"tobialbertino/portfolio-be/config"
	"tobialbertino/portfolio-be/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	var (
		cfg *config.Config
		err error
	)

	// Load config
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Printf("error loading config: %s", err)
	}

	// Use default logger
	file, err := os.OpenFile("./info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	log.SetOutput(file)

	// initiate framework
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.CustomErrorHandler,
	})
	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	// set modules & app Router
	appRouter.InitRouter(app)

	if cfg.Server.Port == "" {
		log.Println("Port tidak ditemukan")
	}
	app.Listen(":" + cfg.Server.Port)
}
