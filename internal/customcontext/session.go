package customcontext

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vndg-rdmt/authpod/internal/entity"
)

type ContextKey string

const (
	KeySession ContextKey = "session"
)

func SetSessionFiber(c *fiber.Ctx, sess *entity.Session) {
	c.SetUserContext(context.WithValue(c.UserContext(), KeySession, sess))
}

func GetSessionFiber(c *fiber.Ctx) *entity.Session {
	if res, ok := c.UserContext().Value(KeySession).(*entity.Session); ok {
		return res
	}
	return nil
}
