package telegram

import (
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"path"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendReplay(message string, video bool) bool {
	log.I("Uploading archived stream to Telegram")
	var file string
	if video {
		file = path.Join(config.Get().CachePath, globals.VideoFilename)
	} else {
		file = path.Join(config.Get().CachePath, globals.ScreenshotFilename)
	}
	_, err := client.SendMedia(config.Get().ReplaysChannelId, file, &tg.MediaOptions{
		FileName:  path.Base(file),
		Caption:   message,
		ParseMode: "markdown",
	})
	if err != nil {
		log.E("Uploading to Telegram failed!")
		log.E(err.Error())
	} else {
		log.I("Video successfully uploaded to Telegram")
	}

	return err == nil
}
