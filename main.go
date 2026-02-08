package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/internal/api"
	"github.com/wynyga/gotoko/internal/config"
	"github.com/wynyga/gotoko/internal/connection"
	"github.com/wynyga/gotoko/internal/repository"
	"github.com/wynyga/gotoko/internal/service"
)

func main() {
	fmt.Println("1. Aplikasi dimulai...")

	cnf := config.Get()
	// Cek apakah config terbaca
	fmt.Printf("2. Mencoba koneksi ke DB: %s di %s\n", cnf.Database.Name, cnf.Database.Host)

	dbConnection := connection.GetDatabase(cnf.Database)
	fmt.Println("3. Database Terhubung!")

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService)

	fmt.Printf("4. Server akan berjalan di port: %s\n", cnf.Server.Port)
	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
