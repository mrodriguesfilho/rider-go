package router

import (
	"context"
	"net/http"
	"rider-go/api/handlers"
	"rider-go/internal/logger"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewChiRouter(signUpHandler *handlers.SignUpHandler, getAccountHandler *handlers.GetAccountHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/account", func(r chi.Router) {
		r.Post("/signup", signUpHandler.Handle)
		r.Get("/", getAccountHandler.Handle)
	})
	return r
}

func StartServer(lc fx.Lifecycle, logger logger.CustomLogger, router *chi.Mux) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Started server")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping server")
			return server.Shutdown(ctx)
		},
	})
}
