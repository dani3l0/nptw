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
		file = path.Join(config.Get().CachePath, config.Get().VideoFilename)
	} else {
		file = path.Join(config.Get().CachePath, config.Get().ScreenshotFilename)
	}
	_, err := client.SendMedia(config.Get().ChannelId, file, &tg.MediaOptions{
		FileName:  path.Base(file),
		Caption:   message,
		ParseMode: "markdown",
	})
	if err != nil {
		utils.Err("Uploading to Telegram failed!")
		utils.Err(err.Error())
	} else {
		utils.Log("Video successfully uploaded to Telegram")
	}

	return err == nil
}
