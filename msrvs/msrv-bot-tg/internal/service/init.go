package service

import (
	"context"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/repository/postgres"
	"projectX/msrvs/msrv-bot-tg/internal/repository/stack"
	"projectX/pkg/model"
)

type IServiceGet interface {
	GetEvent(ctx context.Context) model.Message
}
type IServiceSet interface {
	SetEvent(ctx context.Context, req []byte)
}

type IService interface {
	GetUsersByGroup(ctx context.Context, group string) ([]int64, error)
	AddUser(ctx context.Context, user string, chatID int64) error
	GetUser(ctx context.Context, user string) (chatID int64, err error)

	IServiceGet
	IServiceSet
}

func InitService(cnf *config.Config, stack stack.IStackEvent, db postgres.IRepository) IService {
	return &service{db: db, stack: stack}
}

type service struct {
	db    postgres.IRepository
	stack stack.IStackEvent
}
