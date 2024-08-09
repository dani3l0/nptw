package telegram

import (
	"nptw/config"
	"nptw/utils"
	"path"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendReplay(message string, video bool) bool {
	utils.Log("Uploading archived stream to Telegram")
	var file string
	if video {
		file = "replay.mp4"
	} else {
		file = path.Join(config.Get().CachePath, "screenshot.jpg")
	}
	_, err := client.SendMedia(config.Get().ChannelId, file, &tg.MediaOptions{
		FileName:  path.Base(file),
		Caption:   message,
		ParseMode: "markdown",
	})
	if err != nil {
		utils.Err("Uploading to Telegram failed!")
		utils.Err(err.Error())
	}

	return err == nil
}
