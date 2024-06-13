package main

import (
	"log/slog"
	"projectX/msrvs/msrv-produser/config"
	"projectX/msrvs/msrv-produser/internal/api"
	"projectX/msrvs/msrv-produser/internal/broker"
	"projectX/msrvs/msrv-produser/internal/repository"
	"projectX/msrvs/msrv-produser/internal/service"
)

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	slog.Info("load config success")

	rep := repository.InitRepository(cnf)
	br := broker.InitBroker(cnf.BrokerHost)
	defer br.Close()

	srv := service.InitService(rep, br)
	app := api.InitRestApi(srv)

	app.Start(cnf.RestHost)
}
