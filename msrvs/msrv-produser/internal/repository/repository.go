package repository

import "projectX/msrvs/msrv-produser/config"

type IRepository interface {
	Set()
	Get()
}

func InitRestApi(cnf *config.Config) IRepository {
	return &mongoDB{}
}

type mongoDB struct {
}

func (h *mongoDB) Set() {

}

func (h *mongoDB) Get() {

}
