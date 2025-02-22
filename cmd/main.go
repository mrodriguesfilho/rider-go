package main

import (
	"context"
	"rider-go/api/handlers"
	"rider-go/internal/application/usecase"
	"rider-go/internal/domain/entity"
	inmemory "rider-go/internal/infra/database/InMemory"
	"rider-go/internal/infra/database/repository"
	"rider-go/internal/infra/logger"
	"rider-go/internal/infra/otel"
	"rider-go/internal/infra/router"

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

func NewAccountRepositoryWithDb() repository.AccountRepository {
	return inmemory.NewAccountRepository(make([]entity.Account, 0))
}
