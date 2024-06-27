package main

import (
	"context"
	"os"
	"os/signal"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/api/bot"
	"projectX/msrvs/msrv-bot-tg/internal/api/broker"
	"projectX/msrvs/msrv-bot-tg/internal/repository/cacheEvent"
	"projectX/msrvs/msrv-bot-tg/internal/repository/cacheEvent/memorySlice"
	"projectX/msrvs/msrv-bot-tg/internal/repository/cacheEvent/redis"
	stack2 "projectX/msrvs/msrv-bot-tg/internal/repository/cacheEvent/stack"
	"projectX/msrvs/msrv-bot-tg/internal/repository/postgres"
	"projectX/msrvs/msrv-bot-tg/internal/service"
	"syscall"
)

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	rep := postgres.InitRepository(cnf)

	var stack cacheEvent.IStackEvent
	if cnf.Redis.Work {
		stack = redis.InitRedis(cnf.Redis)
		defer stack.Close()
	} else {
		stack = memorySlice.InitCache()
	}
	srv := service.InitService(cnf, stack, rep)
	br := broker.InitBroker(cnf.Broker.Host, srv)
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
