package service

import (
	"projectX/msrvs/msrv-produser/config"
	"projectX/msrvs/msrv-produser/internal/repository"
)

type IService interface {
	Set()
	Get()
}

func InitRestApi(cnf *config.Config, db *repository.IRepository) IService {
	return &service{}
}

type service struct {
}

func (h *service) Set() {

}

func (h *service) Get() {

}
