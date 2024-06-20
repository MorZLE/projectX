package bot

import (
	"fmt"
	bot "gopkg.in/telebot.v3"
	"log/slog"
	"projectX/pkg/model"
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
	id := h.getRecipient(msg.Group)
	_, err := h.bot.Send(&MyRecipient{user: id}, msg.Body)
	if err != nil {
		slog.Error("error send", err, op)
	}
}

func (h *Bot) HStart(c bot.Context) error {
	const op = "bot.HStart"
	h.addRecipient(c)
	h.bot.Send(c.Chat(), "Привет")
	return nil
}

func (h *Bot) addRecipient(c bot.Context) {
	id := c.Chat().ID
	user := c.Sender().Username
	h.userToChat[user] = id

}

func (h *Bot) getRecipient(user string) (id int64) {
	return h.userToChat[user]
}
