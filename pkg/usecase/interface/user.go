package services

import (
	"context"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
