package main

import (
	"context"
	"os"
	"os/signal"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/api"
	"projectX/msrvs/msrv-bot-tg/internal/bot"
	"projectX/msrvs/msrv-bot-tg/internal/repository"
	"projectX/msrvs/msrv-bot-tg/internal/service"
	"syscall"
)

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	rep := repository.InitRepository(cnf)
	srv := service.InitService(cnf, rep)
	br := api.InitBroker(cnf.Broker.Host, srv)
	defer br.Close()

	iBot, err1 := bot.InitBot(cnf, &srv)
	if err1 != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go iBot.Start(ctx)
	go br.WatchEvents()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
}
