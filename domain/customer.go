package domain

import (
	"context"
	"database/sql"

	"github.com/wynyga/gotoko/dto"
)

type Customer struct {
	ID        string       `db:"id"`
	Code      string       `db:"code"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, c *Customer) error
	Update(ctx context.Context, c *Customer) error
	Delete(ctx context.Context, id string) error
}

type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) error
	Update(ctxx context.Context, req dto.UpdateCustomerRequest) error
	Delete(ctx context.Context, id string) error                   //Step 1 domain 2 service 3 repository 4. Api
	Show(ctx context.Context, id string) (dto.CustomerData, error) //Jangan lupa nama methodnya harus besar agar bisa diakses oleh package lain
}
