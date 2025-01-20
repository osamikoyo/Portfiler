package app

import (
	"github.com/osamikoyo/portfiler/internal/transport/handler"
	"github.com/osamikoyo/portfiler/pkg/loger"
)

type App struct{
	logger loger.Logger
	handler handler.Handler
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
	}
}