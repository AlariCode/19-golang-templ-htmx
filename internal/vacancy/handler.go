package vacancy

import (
	"alaricode/go-fiber/pkg/tadapter"
	"alaricode/go-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	email := c.FormValue("email")
	var component templ.Component
	if email == "" {
		component = components.Notification("Не задан email", components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
