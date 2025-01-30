package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}
		userEmail := ""
		if email, ok := sess.Get("email").(string); ok {
			userEmail = email
		}
		c.Locals("email", userEmail)
		return c.Next()
	}
}
