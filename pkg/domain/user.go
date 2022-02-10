package domain

import (
	"context"
)

type Users struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]Users, error)
	FindByID(ctx context.Context, id uint) (Users, error)
	Save(ctx context.Context, user Users) (Users, error)
	Delete(ctx context.Context, user Users) error
}
