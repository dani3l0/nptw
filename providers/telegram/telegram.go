package telegram

import (
	"fmt"
	"nptw/config"
	"nptw/utils/log"
	"os"

	tg "github.com/amarnathcjd/gogram/telegram"
)

var client *tg.Client
var myID int64

func Init() {
	client, _ = tg.NewClient(tg.ClientConfig{
		AppID:   int32(config.Get().TelegramApiId),
		AppHash: config.Get().TelegramApiHash,
	})

	err := client.ConnectBot(config.Get().TelegramBotToken)
	if err != nil {
		log.E("Couldn't connect to Telegram!")
		log.E(err.Error())
		os.Exit(1)
	}

	myID = client.Me().ID
	log.I(fmt.Sprint("Bot's ID is ", myID))

	if config.Get().RespondToUserMessagesEnabled {
		client.AddMessageHandler(tg.OnNewMessage, func(message *tg.NewMessage) error {
			ReplyToUser(message)
			return nil
		})
	}

	log.I("Bot ready and running")
}
