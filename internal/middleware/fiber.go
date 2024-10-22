package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vndg-rdmt/authpod/internal/auth"
	"github.com/vndg-rdmt/authpod/internal/customcontext"
	"github.com/vndg-rdmt/authpod/internal/entity"
)

const (
	HeaderAuthorization = fiber.HeaderAuthorization
)

func FiberAuthentication(
	methods auth.Authentication,
	timeout time.Duration,
) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(c.UserContext(), timeout)
		defer cancel()

		// get credetials
		credentials := c.Get(HeaderAuthorization, "")
		if credentials == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		authArgs := strings.SplitN(credentials, "", 2)
		if authArgs == nil || len(authArgs) < 2 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// authenticate user
		var user entity.User
		ok, err := methods.Authenticate(ctx, &user, authArgs[0], authArgs[1])
		if err != nil {
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		session := entity.Session{
			UserID: user.Id,
		}

		// pass session to futher handlers
		customcontext.SetSessionFiber(c, &session)
		return c.Next()
	}
}
