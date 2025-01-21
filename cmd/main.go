package main

import (
	"alaricode/go-fiber/config"
	"alaricode/go-fiber/internal/home"
	"alaricode/go-fiber/internal/vacancy"
	"alaricode/go-fiber/pkg/database"
	"alaricode/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	home.NewHandler(app, customLogger)
	vacancy.NewHandler(app, customLogger)

	app.Listen(":3000")
}
