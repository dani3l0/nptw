package telegram

import (
	"fmt"
	"nptw/config"
	"nptw/utils/log"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail string, title string) {
	log.I("Sending notification about started stream to Telegram")
	log.V("Thumbnail: ", thumbnail)

	text := fmt.Sprintf("**🇵🇱 Rozpoczął się Żywiec! 🇵🇱**\n\n__🐺 %s 🦎__\n\n🔗 Link do DLive: https://dlive.tv/nptvpl", title)

	_, err := client.SendMedia(config.Get().NotificationsChannelId, thumbnail, &tg.MediaOptions{
		Caption:   text,
		ParseMode: "markdown",
	})
	if err != nil {
		log.E("Couldn't send notification!")
		log.E(err.Error())
	}
}
