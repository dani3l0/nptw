package telegram

import (
	"nptw/config"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var client *tg.Client

func Init() {
	client, _ = tg.NewClient(tg.ClientConfig{
		AppID:   int32(config.Get().TelegramApiId),
		AppHash: config.Get().TelegramApiHash,
	})

	client.ConnectBot(config.Get().TelegramBotToken)

	client.AddMessageHandler(tg.OnNewMessage, func(message *tg.NewMessage) error {
		Chat(message)
		return nil
	})
}
