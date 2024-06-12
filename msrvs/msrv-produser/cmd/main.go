package main

import (
	"projectX/msrvs/msrv-produser/config"
	"projectX/msrvs/msrv-produser/internal/api"
	"projectX/msrvs/msrv-produser/internal/repository"
	"projectX/msrvs/msrv-produser/internal/service"
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
