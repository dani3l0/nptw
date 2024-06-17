package telegram

import tg "github.com/amarnathcjd/gogram/telegram"

func Chat(message *tg.NewMessage) {
	if !message.IsPrivate() {
		return
	}

	if message.Text() == "/start" {
		message.Reply("Hello!")
	}
}
