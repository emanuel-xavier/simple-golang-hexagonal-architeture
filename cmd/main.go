package main

import (
	"github.com/emanuel-xavier/hexagonal-architerure/configs"
	adapters "github.com/emanuel-xavier/hexagonal-architerure/internal/adapters/postgres"
	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/usecase"
	httpHandler "github.com/emanuel-xavier/hexagonal-architerure/internal/handler/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	psql, err := adapters.NewPostgresUserRepository()
	if err != nil {
		panic(err)
	}

	userUCase := usecase.NewUserUseCase(psql)

	handler := httpHandler.NewHandler(userUCase)
	handler.Handle()

}
