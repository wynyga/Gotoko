package main

import (
	"fmt"
	"log"
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/wynyga/gotoko/dto"
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

	//MiddleWare
	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("Endpoint perlu token"))
		},
	})

	//Auth
	userRepository := repository.NewUser(dbConnection)
	authService := service.NewAuth(cnf, userRepository)
	api.NewAuth(app, authService)

	//Customer
	customerRepository := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepository)
	api.NewCustomer(app, customerService, jwtMidd)

	fmt.Printf("4. Server akan berjalan di port: %s\n", cnf.Server.Port)
	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
