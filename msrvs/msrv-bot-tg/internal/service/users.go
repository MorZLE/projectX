package service

import (
	"context"
	"errors"
	"log/slog"
)

func (s *service) GetUsersByGroup(ctx context.Context, group string) ([]int64, error) {
	switch group {
	case "all":
		return s.db.GetAllUsers(ctx)
	default:
		slog.Error("not found group", group)
		//return s.db.GetUsersByGroup(ctx, group)
		return nil, errors.New("not found group")
	}

}

func (s *service) AddUser(ctx context.Context, user string, chatID int64) error {
	return s.db.AddUser(ctx, user, chatID)
}

func (s *service) GetUser(ctx context.Context, user string) (chatID int64, err error) {
	return s.db.GetUser(ctx, user)
}
