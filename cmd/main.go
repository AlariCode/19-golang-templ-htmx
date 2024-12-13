package main

import (
	"alaricode/go-fiber/config"
	"alaricode/go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
