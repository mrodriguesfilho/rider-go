package router

import (
	"context"
	"log"
	"net/http"
	"rider-go/api/handlers"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewChiRouter(signUpHandler *handlers.SignUpHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/account", func(r chi.Router) {
		r.Post("/signup", signUpHandler.Handle)
	})
	return r
}

func StartServer(lc fx.Lifecycle, router *chi.Mux) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Started server")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("Stopping server")
			return server.Shutdown(ctx)
		},
	})
}
