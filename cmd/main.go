package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/osamikoyo/portfiler/internal/app"
	"github.com/osamikoyo/portfiler/pkg/loger"
)

func main() {
	ctx,cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := app.Init()
	if err := app.Run(ctx);err != nil{
		loger.New().Error().Err(err)
	}
}