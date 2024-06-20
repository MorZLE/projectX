package service

import (
	"context"
	"encoding/json"
	"log/slog"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/repository"
	"projectX/pkg/model"
)

type IServiceGet interface {
	Get(ctx context.Context) model.Message
}
type IServiceSet interface {
	Set(ctx context.Context, req []byte)
}

type IService interface {
	IServiceGet
	IServiceSet
}

func InitService(cnf *config.Config, db repository.IRepository) IService {
	srv := service{db: db}
	return &srv
}

type service struct {
	db repository.IRepository
}

func (s *service) Set(ctx context.Context, req []byte) {
	const op = "service.Set"

	msg := model.Message{}

	err := json.Unmarshal(req, &msg)
	if err != nil {
		slog.Error("error unmarshal", err, string(req), op)
		return
	}

	s.db.Set(ctx, msg)
}

func (s *service) Get(ctx context.Context) model.Message {
	msg, err := s.db.Get()
	if err != nil {
		msg.Error = err
		return msg
	}
	return msg
}
