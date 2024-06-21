package telegram

import (
	"nptw/config"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail interface{}, message string) {
	client.SendMedia(config.Get().ChannelId, thumbnail, &tg.MediaOptions{
		Caption:   message,
		ParseMode: "markdown",
	})
}
