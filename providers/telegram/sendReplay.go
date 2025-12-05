package telegram

import (
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"path"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendReplay(message string, thumb any) bool {
	log.I("Uploading archived stream to Telegram")
	file := path.Join(config.Get().CachePath, globals.VideoFilename)
	_, err := client.SendMedia(config.Get().ReplaysChannelId, file, &tg.MediaOptions{
		FileName:  path.Base(file),
		Caption:   message,
		ParseMode: "markdown",
		Thumb:     thumb,
		ProgressCallback: func(pi *tg.ProgressInfo) {
			ProgressUp.UploadSize = int(pi.TotalSize) / 1000 / 1000
			ProgressUp.UploadedBytes = pi.Percentage * float64(pi.TotalSize) / 100 / 1000 / 1000
		},
		ProgressInterval: 5,
	})
	ProgressUp.Stage = 3
	ProgressUp.Done = true
	if err != nil {
		log.E("Uploading to Telegram failed!")
		log.E(err.Error())
	} else {
		log.I("Video successfully uploaded to Telegram")
	}

	return err == nil
}
