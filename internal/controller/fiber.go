package controller

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vndg-rdmt/authpod/internal/controller/request"
	"github.com/vndg-rdmt/authpod/internal/controller/response"
	"github.com/vndg-rdmt/authpod/internal/service"
)

func NewFiber(srv service.Service) *Fiber {
	return &Fiber{
		service: srv,
	}
}

type Fiber struct {
	service service.Service
}

func (f *Fiber) SignIn(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()

	var request request.SignIn
	var err error

	dec := json.NewDecoder(bytes.NewBuffer(c.Body()))
	dec.DisallowUnknownFields()
	if err = dec.Decode(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	sess, err := f.service.SignIn(ctx, request.Login, request.Password)
	if err != nil {
		if err == service.ErrNotFound {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusServiceUnavailable)
	}

	return c.Status(fiber.StatusOK).JSON(response.SignIn{
		SessionId: sess,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})

}

func (f *Fiber) Ping(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()

	header := c.Get(fiber.HeaderAuthorization, "")
	if header == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	res, err := f.service.Ping(ctx, header)
	if err != nil {
		if err == service.ErrNotFound {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusServiceUnavailable)
	}

	return c.Status(fiber.StatusOK).JSON(response.Ping{
		UserId: res,
	})
}

func (f *Fiber) IssueToken(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()

	var userId int64
	var err error

	if rawUserId := c.Params("user_id", ""); rawUserId == "" {
		return c.SendStatus(fiber.StatusNotFound)
	} else if userId, err = strconv.ParseInt(rawUserId, 10, 0); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var token service.Token
	if err = f.service.IssueToken(ctx, &token, userId); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

func (f *Fiber) CheckToken(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()

	var token string
	var userid int64
	var err error

	if token = c.Params("token", ""); token == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if userid, err = f.service.CheckToken(ctx, token); err != nil {
		if err == service.ErrNotFound {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusServiceUnavailable)
	}

	return c.Status(fiber.StatusOK).JSON(response.Ping{
		UserId: userid,
	})
}
