//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	db "github.com/thnkrn/go-gin-clean-arch/pkg/db"
	http "github.com/thnkrn/go-gin-clean-arch/pkg/http"
	handler "github.com/thnkrn/go-gin-clean-arch/pkg/http/handler"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase"
	repository "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/repository"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
