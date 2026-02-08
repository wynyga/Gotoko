package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/internal/config"
	"github.com/wynyga/gotoko/internal/connection"
	"github.com/wynyga/gotoko/internal/repository"
	"github.com/wynyga/gotoko/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)

	app.Get("/developers", developers)
	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
