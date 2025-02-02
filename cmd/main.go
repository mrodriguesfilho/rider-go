package main

import (
	"rider-go/api/handlers"
	router "rider-go/api/router/chi"
	"rider-go/internal/entity"
	"rider-go/internal/infra/database"
	"rider-go/internal/usecase"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(NewAccountRepositoryWithDb),
		fx.Provide(usecase.NewSignUpUseCase),
		fx.Provide(handlers.NewSignUpHandler),
		fx.Provide(usecase.NewGetAccountUseCase),
		fx.Provide(handlers.NewGetAccountHandler),
		fx.Provide(router.NewChiRouter),
		fx.Invoke(router.StartServer),
	)

	app.Run()
}

func NewAccountRepositoryWithDb() *database.AccountRepositoryInMemory {
	return database.NewAccountRepository(make([]entity.Account, 0))
}
