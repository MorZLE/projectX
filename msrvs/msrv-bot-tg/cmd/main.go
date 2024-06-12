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

	rep := repository.InitRestApi(cnf)
	srv := service.InitRestApi(cnf, &rep)
	app := api.InitRestApi(cnf, &srv)

	go app.Start("")
}
