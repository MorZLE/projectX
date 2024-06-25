package service

import (
	"context"
	"errors"
	"log/slog"
)

func (s *service) GetUsersByGroup(ctx context.Context, group string) ([]int64, error) {
	const op = "service.GetUsersByGroup"
	switch group {
	case "all":
		return s.db.GetAllUsers(ctx)
	default:
		users, err := s.db.GetUserByGroup(ctx, group)
		if err != nil {
			slog.Error("error get users", err, group, op)
			return nil, err
		}
		if len(users) == 0 {
			slog.Info("users not found", group, op)
			return nil, errors.New("users not found")
		}
		return users, nil
	}
}

func (s *service) AddUser(ctx context.Context, user string, chatID int64) error {
	if user == "" {
		return errors.New("user is empty")
	}
	if _, ok := s.hiAdmin[user]; ok {
		slog.Info("add user admin", user)
		return s.db.AddUser(ctx, user, chatID, admin)
	}

	return s.db.AddUser(ctx, user, chatID, defaultGroup)
}

func (s *service) GetUser(ctx context.Context, user string) (chatID int64, err error) {
	return s.db.GetUser(ctx, user)
}
