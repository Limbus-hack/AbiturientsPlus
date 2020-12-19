package main

import (
	"github.com/code7unner/vk-scrapper/config"
	"github.com/code7unner/vk-scrapper/internal/api/handler"
	"github.com/code7unner/vk-scrapper/internal/api/repository"
	"github.com/code7unner/vk-scrapper/internal/app"
	"github.com/code7unner/vk-scrapper/internal/db"
	"github.com/code7unner/vk-scrapper/internal/interrupt"
	"github.com/code7unner/vk-scrapper/internal/server"
	"github.com/code7unner/vk-scrapper/logger"
)

func main() {
	ctx, done := interrupt.Context()
	defer done()

	// Init config
	conf := config.GetCommonEnvConfigs()

	// Init logger
	log := logger.NewLogger(logger.NewPreparedStdoutCore(conf.LogLevel))
	logger.SetDefaultLogger(log.Named("api"))

	// Init Database
	database, err := db.New(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	// Init repository
	r := repository.New(database.Conn, log, conf)

	// Init app logic
	a := app.New(log, conf, r, ctx)

	// Init handler
	h := handler.New(a)

	// Init Server
	srv, err := server.New(conf.ServerPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("listening on :%s", conf.ServerPort)

	// Start server
	if err := srv.ServeHTTPHandler(ctx, h.Mux); err != nil {
		log.Fatal(err)
	}
}
