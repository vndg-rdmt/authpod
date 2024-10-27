package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vndg-rdmt/authpod/internal/controller"
)

func NewHttp(ctr *controller.Fiber, host string) error {
	app := fiber.New()
	app.Post("/api/sign-in", ctr.SignIn)
	app.Get("/api/ping", ctr.Ping)
	app.Put("/api/tokens/:user_id", ctr.IssueToken)
	app.Get("/api/tokens/:token", ctr.CheckToken)

	return app.Listen(host)
}
