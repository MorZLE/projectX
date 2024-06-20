package bot

import (
	"context"
	bot "gopkg.in/telebot.v3"
	"log/slog"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/service"
	"projectX/msrvs/msrv-bot-tg/internal/worker"
	"projectX/pkg/model"
	"time"
)

type IBot interface {
	Start(ctx context.Context)
	SendEvent(msg model.Message)
}

func InitBot(cnf *config.Config, srv *service.IService) (IBot, error) {
	pref := bot.Settings{
		Token:  cnf.Bot.Token,
		Poller: &bot.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := bot.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &Bot{bot: b, srv: srv, watcher: worker.InitWatcher(srv)}, nil
}

type Bot struct {
	bot     *bot.Bot
	watcher worker.IWatcher
	srv     *service.IService
}

func (h *Bot) Start(ctx context.Context) {
	const op = "bot.Start"

	go h.bot.Start()
	h.initHandler()
	slog.Info("Bot started")

	out := make(chan model.Message)
	defer close(out)

	go h.watcher.WatcherEvent(ctx, out)
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-out:
			if !ok {
				return
			}
			slog.Info("msg", msg, op)
			h.SendEvent(msg)
		}
	}
}

func (h *Bot) initHandler() {
	h.bot.Handle("/start", h.HStart)
}
