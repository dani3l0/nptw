package telegram

import (
	"nptw/config"
	"nptw/utils"
	"os"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var client *tg.Client

func Init() {
	client, _ = tg.NewClient(tg.ClientConfig{
		AppID:   int32(config.Get().TelegramApiId),
		AppHash: config.Get().TelegramApiHash,
	})

	err := client.ConnectBot(config.Get().TelegramBotToken)
	if err != nil {
		utils.Err("Couldn't connect to Telegram!")
		utils.Err(err.Error())
		os.Exit(1)
	}

	if config.Get().RespondEnabled {
		client.AddMessageHandler(tg.OnNewMessage, func(message *tg.NewMessage) error {
			ReplyToUser(message)
			return nil
		})
	}
}
