package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/domain"
)

type customerAPI struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService){
	ca := customerAPI{
		customerService: customerService,
	}
}

func (ca customerAPI)
