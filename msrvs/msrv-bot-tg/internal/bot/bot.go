package bot

import (
	bot "gopkg.in/telebot.v3"
	"log/slog"
	"projectX/pkg/model"
)

type MyRecipient struct {
	user string
}

func (r *MyRecipient) Recipient() string {
	return r.user
}

func (h *Bot) SendEvent(msg model.Message) {
	const op = "bot.SendEvent"

	if msg.Error != nil {
		return
	}
	_, err := h.bot.Send(&MyRecipient{user: msg.Group}, msg.Body)
	if err != nil {
		slog.Error("error send", err, op)
	}
}

func (h *Bot) HStart(c bot.Context) error {
	h.bot.Send(c.Chat(), "Привет")
	return nil
}
