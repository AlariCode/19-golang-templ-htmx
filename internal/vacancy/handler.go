package vacancy

import (
	"alaricode/go-fiber/pkg/tadapter"
	"alaricode/go-fiber/pkg/validator"
	"alaricode/go-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
