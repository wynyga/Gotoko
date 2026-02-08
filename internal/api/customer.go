package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/domain"
	"github.com/wynyga/gotoko/dto"
	"github.com/wynyga/gotoko/internal/util"
)

type customerAPI struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := customerAPI{
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
}

func (ca customerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ca customerAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var reg dto.CreateCustomerRequest
	if err := ctx.BodyParser(&reg); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(reg)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}
	err := ca.customerService.Create(c, reg)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(" "))

}
