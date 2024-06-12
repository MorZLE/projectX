package api

import (
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/service"
)

type IBot interface {
	Start(addr string)
}

func InitRestApi(cnf *config.Config, srv *service.IService) IBot {
	return &Bot{}
}

type Bot struct {
}

func (h *Bot) Start(string) {

}
