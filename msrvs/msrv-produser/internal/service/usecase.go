package service

import (
	"context"
	"encoding/json"
	"projectX/msrvs/msrv-produser/internal/broker"
	"projectX/msrvs/msrv-produser/internal/repository"
	"projectX/msrvs/pkg/cerrors"
	model2 "projectX/msrvs/pkg/model"
)

type IService interface {
	Set(ctx *context.Context, req *model2.UserReq) error
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

func (h *Service) Set(ctx *context.Context, req *model2.UserReq) error {
	if req.Username == "" {
		return cerrors.ErrUserNil
	}
	if req.Body == "" {
		return cerrors.ErrBodyNil
	}

	var msg model2.Message
	msg.Topic = "user"
	msg.Body = req.Body
	msg.Group = req.Username
	msg.Status = "sent"

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = h.br.Send(ctx, "event", body)
	if err != nil {
		return err
	}
	return nil
}

func (h *Service) Get() {

}
