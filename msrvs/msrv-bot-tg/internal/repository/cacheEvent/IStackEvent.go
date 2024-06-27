package cacheEvent

import (
	"context"
	"projectX/msrvs/pkg/model"
)

type IStackEvent interface {
	Set(ctx context.Context, msg model.Message)
	Get() (msg model.Message, err error)
	Close() error
	Ping(ctx context.Context) error
}
