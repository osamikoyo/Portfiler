package app

import (
	"context"
	"net/http"

	"github.com/osamikoyo/portfiler/internal/transport/handler"
	"github.com/osamikoyo/portfiler/pkg/loger"
)

type App struct{
	logger loger.Logger
	handler handler.Handler
	server *http.Server
}

func Init() App {
	loger := loger.New()
	handler, err := handler.New()
	if err != nil{
		loger.Error().Err(err)
	}
	return App{
		handler: handler,
		logger: loger,
		server: &http.Server{
			Addr: "localhost:8080",
		},
	}
}

func (a *App) Run(ctx context.Context) error {
	go func ()  {
		<- ctx.Done()
		a.logger.Info().Msg("Server shutdown")
		a.server.Shutdown(ctx)
	}()

	mux := http.NewServeMux()

	a.handler.RegisterRouter(mux)

	a.logger.Info().Msg("Server started on :8080")
	return a.server.ListenAndServe()
}