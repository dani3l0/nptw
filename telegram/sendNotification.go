package telegram

import (
	"nptw/config"
	"nptw/utils"
	"path"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail interface{}, message string) {
	utils.Log("Sending notification about started stream to Telegram")
	client.SendMedia(config.Get().ChannelId, thumbnail, &tg.MediaOptions{
		Caption:   message,
		ParseMode: "markdown",
	})
}

func SendReplay(message string) {
	utils.Log("Sending notification about archived stream to Telegram")
	timestamp := time.Now()
	client.SendMedia(config.Get().ChannelId, path.Join(config.Get().CachePath, "outputto.mp4"), &tg.MediaOptions{
		FileName:  "nptv-" + timestamp.String() + ".mp4",
		Caption:   message,
		ParseMode: "markdown",
	})
}
