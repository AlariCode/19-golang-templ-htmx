package main

import (
	"alaricode/go-fiber/config"
	"alaricode/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	app := fiber.New()
	log.SetLevel(log.Level(logConfig.Level))
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
