package service

import (
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/repository"
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
