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

type bookApi struct {
	bookService domain.BookService
}

func NewBook(app *fiber.App, bookService domain.BookService, authMid fiber.Handler) {

	ba := bookApi{
		bookService: bookService,
	}
	app.Get("/books", authMid, ba.Index)
	app.Post("/books", authMid, ba.Create)
	app.Get("/books/:id", authMid, ba.Show)
	app.Put("/books/:id", authMid, ba.Update)
	app.Delete("/books/:id", authMid, ba.Delete)

}

func (ba bookApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ba.bookService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ba bookApi) Create(ctx *fiber.Ctx) error {
	//Timer
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	//Cek Error
	var req dto.CreateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	//Create Prosess
	err := ba.bookService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess("Book created successfully"))
}

func (ba bookApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	data, err := ba.bookService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(data))
}

func (ba bookApi) Update(ctx *fiber.Ctx) error {
	//Timer
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	//Cek Error
	var req dto.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	//Update Proses
	req.Id = ctx.Params("id") //Ambil dari DTO
	err := ba.bookService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess("Book updated successfully"))
}

func (ba bookApi) Delete(ctx *fiber.Ctx) error {
	//Timer
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	//Cek ID dan Delete Proses
	id := ctx.Params("id")
	err := ba.bookService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess("Book deleted successfully"))
}
