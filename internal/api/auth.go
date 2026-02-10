package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/domain"
	"github.com/wynyga/gotoko/dto"
)

type authAPI struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	aa := authAPI{
		authService: authService,
	}
	app.Post("/auth", aa.Login)
}

func (aa authAPI) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	res, err := aa.authService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(res))
}
