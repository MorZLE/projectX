package bot

import (
	"context"
	"fmt"
	bot "gopkg.in/telebot.v3"
	"log/slog"
	"projectX/msrvs/pkg/model"
)

type MyRecipient struct {
	user int64
}

func (r *MyRecipient) Recipient() string {
	return fmt.Sprint(r.user)
}

func (h *Bot) SendEvent(msg model.Message) {
	const op = "bot.SendEvent"
	if msg.Error != nil {
		return
	}

	users, err := h.srv.GetUsersByGroup(context.Background(), msg.Group)
	if err != nil {
		slog.Error("error send", err, op)
		return
	}
	h.send(msg, users)
}

func (h *Bot) send(msg model.Message, users []int64) {
	const op = "bot.SendEvent"
	if msg.Error != nil {
		return
	}
	for i := 0; i < len(users); i++ {
		_, err := h.bot.Send(&MyRecipient{user: users[i]}, msg.Body)
		if err != nil {
			slog.Error("error send", err, op)
		}
	}
}

func (h *Bot) HStart(c bot.Context) error {
	const op = "bot.HStart"
	h.bot.Send(c.Chat(), "Привет")
	return nil
}

func (h *Bot) HSubscribe(c bot.Context) error {
	const op = "bot.Subscribe"
	err := h.srv.AddUser(context.Background(), c.Sender().Username, c.Chat().ID)
	if err != nil {
		return err
	}
	h.bot.Send(c.Chat(), "Вы подписались на обновление")
	return nil
}
