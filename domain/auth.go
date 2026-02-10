package domain

import (
	"context"

	"github.com/wynyga/gotoko/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
