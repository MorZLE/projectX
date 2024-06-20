package repository

import (
	"context"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/pkg/cerrors"
	"projectX/pkg/model"
	"sync"
)

type IRepository interface {
	Set(ctx context.Context, msg model.Message)
	Get() (msg model.Message, err error)
}

func InitRepository(cnf *config.Config) IRepository {

	return &Cache{memory: []model.Message{}}
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
