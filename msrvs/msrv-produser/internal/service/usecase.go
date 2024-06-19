package service

import (
	"context"
	"projectX/msrvs/msrv-produser/internal/broker"
	"projectX/msrvs/msrv-produser/internal/repository"
	"projectX/pkg/cerrors"
	"projectX/pkg/model"
)

type IService interface {
	Set(ctx *context.Context, req *model.UserReq) error
	Get()
}

func InitService(db repository.IRepository, br broker.IBroker) IService {
	return &Service{
		db: db,
		br: br,
	}
}

type Service struct {
	db repository.IRepository
	br broker.IBroker
}

func (h *Service) Set(ctx *context.Context, req *model.UserReq) error {
	if req.Username == "" {
		return cerrors.ErrUserNil
	}
	if req.Body == "" {
		return cerrors.ErrBodyNil
	}
	return h.br.Send(ctx, "test", req.Body)
}

func (h *Service) Get() {

}
