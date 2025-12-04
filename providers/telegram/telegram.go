package telegram

import (
	"fmt"
	"nptw/config"
	"nptw/utils/log"
	"os"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var client *tg.Client
var myID int64

// Initializes Telegram connection
func Init() {
	client, _ = tg.NewClient(tg.ClientConfig{
		AppID:    int32(config.Get().TelegramApiId),
		AppHash:  config.Get().TelegramApiHash,
		LogLevel: tg.WarnLevel,
	})

	err := client.ConnectBot(config.Get().TelegramBotToken)
	if err != nil {
		log.E("Couldn't connect to Telegram!")
		log.E(err.Error())
		os.Exit(1)
	}

	// Get info about running bot
	myID = client.Me().ID
	log.I(fmt.Sprint("Bot's ID is ", myID))

	// Get info about target channels / refresh cache
	initializeChannel(int64(config.Get().NotificationsChannelId), "notifications")
	initializeChannel(int64(config.Get().ReplaysChannelId), "replays")

	// Enable user-bot interaction
	if config.Get().RespondToUserMessagesEnabled {
		client.AddMessageHandler(tg.OnNewMessage, func(message *tg.NewMessage) error {
			ReplyToUser(message)
			return nil
		})
	}
	log.I("Bot ready and running")
}

// Just a helper function that convers true/false into yes/no
func initializeChannel(id int64, channelType string) {
	log.I("Getting info about ", channelType, " channel")
	for {
		channel, err := client.GetChannel(id)
		client.GetSendableChannel(id)
		if !client.IdInCache(id) {
			log.W("Please send a message to ", channelType, " channel")
			time.Sleep(time.Second)
		} else if err == nil {
			postMessages := "no (REQUIRED!)"
			if channel.AdminRights.PostMessages {
				postMessages = "yes"
			}
			caser := cases.Title(language.English)
			log.I("---------- ", caser.String(channelType), " channel info ----------")
			log.I("Title:                ", channel.Title)
			log.I("Can post messages:    ", postMessages)
			log.I("------------------------------------------------")
			break
		} else {
			log.E("Couldn't get info about ", channelType, " channel: ", err.Error())
			break
		}
	}
}
