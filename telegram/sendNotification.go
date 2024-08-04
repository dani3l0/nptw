package telegram

import (
	"nptw/config"
	"path"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail interface{}, message string) {
	client.SendMedia(config.Get().ChannelId, thumbnail, &tg.MediaOptions{
		Caption:   message,
		ParseMode: "markdown",
	})
}

func SendReplay(message string) {
	timestamp := time.Now()
	client.SendMedia(config.Get().ChannelId, path.Join(config.Get().CachePath, "outputto.mp4"), &tg.MediaOptions{
		FileName:  "nptv-" + timestamp.String() + ".mp4",
		Caption:   message,
		ParseMode: "markdown",
	})
}
