package main

import (
	"context"
	"rider-go/api/handlers"
	router "rider-go/api/router/chi"
	"rider-go/internal/entity"
	"rider-go/internal/infra/database"
	"rider-go/internal/infra/logger"
	"rider-go/internal/infra/otel"
	"rider-go/internal/usecase"

	"go.uber.org/fx"
)

func main() {

	tp, err := otel.InitOtel()

	if err != nil {
		panic(err)
	}

	defer func() { _ = tp.Shutdown(context.Background()) }()

	app := fx.New(
		fx.Provide(NewAccountRepositoryWithDb),
		fx.Provide(usecase.NewSignUpUseCase),
		fx.Provide(handlers.NewSignUpHandler),
		fx.Provide(usecase.NewGetAccountUseCase),
		fx.Provide(handlers.NewGetAccountHandler),
		fx.Provide(router.NewChiRouter),
		fx.Provide(logger.NewLogger),
		fx.Invoke(router.StartServer),
	)

	app.Run()
}

func NewAccountRepositoryWithDb() database.AccountRepository {
	return database.NewAccountRepository(make([]entity.Account, 0))
}
