package telegram

import (
	"nptw/config"
	"nptw/utils/log"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail interface{}, message string) {
	log.I("Sending notification about started stream to Telegram")
	_, err := client.SendMedia(config.Get().NotificationsChannelId, thumbnail, &tg.MediaOptions{
		Caption:   message,
		ParseMode: "markdown",
	})
	if err != nil {
		log.E("Couldn't send notification!")
		log.E(err.Error())
	}
}
