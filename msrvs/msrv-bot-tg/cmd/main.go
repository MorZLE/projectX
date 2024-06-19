package main

import (
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/api"
	"projectX/msrvs/msrv-bot-tg/internal/repository"
	"projectX/msrvs/msrv-bot-tg/internal/service"
)

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	rep := repository.InitRepository(cnf)
	srv := service.InitService(cnf, &rep)
	br := api.InitBroker(cnf.BrokerHost, srv)
	defer br.Close()
	br.Get()

}
