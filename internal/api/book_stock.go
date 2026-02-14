package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/domain"
	"github.com/wynyga/gotoko/dto"
	"github.com/wynyga/gotoko/internal/util"
)

type bookStockApi struct {
	bookStockSerice domain.BookStockService
}

func NewBookStockApi(app *fiber.App,
	bookStockService domain.BookStockService, authMidd fiber.Handler) {
	bsa := bookStockApi{
		bookStockSerice: bookStockService,
	}

	app.Post("/book-stocks", authMidd, bsa.Create)
	app.Delete("/book-stocks", authMidd, bsa.Delete)
}

func (ba bookStockApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreatBookStockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validasi gagal", fails))
	}
	err := ba.bookStockSerice.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (ba bookStockApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	codeStr := ctx.Query("code")
	if codeStr == "" {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError("kode stok buku tidak boleh kosong"))
	}
	codes := strings.Split(codeStr, ",")

	err := ba.bookStockSerice.Delete(c, dto.DeleteBookStockRequest{
		Codes: codes,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}
