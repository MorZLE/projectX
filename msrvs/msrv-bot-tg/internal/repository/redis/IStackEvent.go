package redis

import (
	"context"
	"log/slog"
	"projectX/pkg/cerrors"
	"projectX/pkg/model"
)

func (r *RedisDB) Set(ctx context.Context, msg model.Message) {
	res, err := msg.MarshalBinary()
	if err != nil {
		slog.Error("RedisDB.Set marshal err", err)
	}

	err = r.db.RPush(ctx, "messages", res).Err()
	if err != nil {
		slog.Error("RedisDB.Set", err)
	}
}

func (r *RedisDB) Get() (model.Message, error) {
	var user model.Message
	err := r.db.LPop(context.Background(), "messages").Scan(&user)
	if err != nil {
		return user, cerrors.ErrMemoryEmpty
	}

	return user, nil
}
