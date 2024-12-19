package main

import (
	"alaricode/go-fiber/config"
	"alaricode/go-fiber/internal/home"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	app := fiber.New()
	zerolog.SetGlobalLevel(zerolog.Level(logConfig.Level))
	app.Use(fiberzerolog.New())
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
