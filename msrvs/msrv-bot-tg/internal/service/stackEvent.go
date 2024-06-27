package service

import (
	"context"
	"encoding/json"
	"log/slog"
	"projectX/msrvs/pkg/model"
)

func (s *service) SetEvent(ctx context.Context, req []byte) {
	const op = "service.SetEvent"

	msg := model.Message{}

	err := json.Unmarshal(req, &msg)
	if err != nil {
		slog.Error("error unmarshal", err, string(req), op)
		return
	}

	s.stack.Set(ctx, msg)
}

func (s *service) GetEvent(ctx context.Context) model.Message {
	msg, err := s.stack.Get()
	if err != nil {
		msg.Error = err
		return msg
	}
	return msg
}
