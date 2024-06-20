package worker

import (
	"context"
	"errors"
	"log/slog"
	"projectX/msrvs/msrv-bot-tg/internal/service"
	"projectX/pkg/cerrors"
	"projectX/pkg/model"
)

func InitWatcher(srv *service.IService) IWatcher {
	return &Watcher{srv: *srv}
}

type IWatcher interface {
	WatcherEvent(ctx context.Context, out chan<- model.Message)
	CloseWatcher()
}
type Watcher struct {
	srv service.IService
}

func (w *Watcher) CloseWatcher() {
	// TODO: implement
}

func (w *Watcher) WatcherEvent(ctx context.Context, out chan<- model.Message) {
	const op = "Watcher.WatcherEvent"

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg := w.srv.Get(ctx)
			if msg.Error != nil {
				if errors.Is(msg.Error, cerrors.ErrMemoryEmpty) {
					continue
				}
				slog.Error("WatcherEvent", msg.Error, op)
				continue
			}
			select {
			case out <- msg:
			}
		}
	}
}
