package telegram

import (
	"nptw/config"
	"nptw/utils"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendReplay(title string, message string) {
	utils.Log("Uploading archived stream to Telegram: `" + title + "`")
	client.SendMedia(config.Get().ChannelId, "replay.mp4", &tg.MediaOptions{
		FileName:  title + ".mp4",
		Caption:   message,
		ParseMode: "markdown",
	})
}
