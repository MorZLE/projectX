package memorySlice

import (
	"context"
	"log/slog"
	"projectX/msrvs/msrv-bot-tg/internal/repository/cacheEvent"
	"projectX/msrvs/pkg/cerrors"
	"projectX/msrvs/pkg/model"
	"sync"
)

func InitCache() cacheEvent.IStackEvent {
	slog.Info("init memoryCache")
	return &Cache{memory: make([]model.Message, 0)}
}

type Cache struct {
	mt     sync.Mutex
	memory []model.Message
}

func (m *Cache) Set(ctx context.Context, msg model.Message) {
	m.mt.Lock()
	defer m.mt.Unlock()
	m.memory = append(m.memory, msg)
}

func (m *Cache) Get() (model.Message, error) {
	m.mt.Lock()
	defer m.mt.Unlock()
	var msg model.Message

	if len(m.memory) == 0 {
		return msg, cerrors.ErrMemoryEmpty
	}

	msg = m.memory[len(m.memory)-1]
	m.memory = m.memory[:len(m.memory)-1]

	return msg, nil
}

func (m *Cache) Close() error {
	return nil
}

func (m *Cache) Ping(ctx context.Context) error {
	return nil
}
