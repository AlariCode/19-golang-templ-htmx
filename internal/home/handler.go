package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")
	return c.SendString("Error")
}
