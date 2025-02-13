package telegram

import (
	"fmt"
	"math/rand"
	"nptw/config"
	"nptw/utils/log"
	"path/filepath"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail string, title string) {
	log.I("Sending notification about started stream to Telegram")

	// If thumbnail does not exist, fallback to fs or online image
	if thumbnail == "" {
		log.W("Live thumbnail seems to be unavailable")
		files, err := filepath.Glob("images/*")
		if err != nil || len(files) == 0 {
			log.E("Custom thumb dir error: `images` folder seems not to exist or is empty!")
			log.W("Falling back to nptv.pl image")
			thumbnail = "https://nptv.pl/lib/k0fssq/BG1-ko2dg25d.jpg"
		} else {
			idx := rand.Intn(len(files))
			thumbnail = files[idx]
			log.I("Selected local thumbnail `", thumbnail, "`")
		}
	}

	text := fmt.Sprintf("<b>🇵🇱 Rozpoczął się Żywiec! 🇵🇱</b>\n\n<i>🐺 %s 🦎</i>\n\n🔗 Link do DLive: https://dlive.tv/"+config.Get().DliveUsername, title)
	log.V(text)

	_, err := client.SendMedia(config.Get().NotificationsChannelId, thumbnail, &tg.MediaOptions{
		Caption:   text,
		ParseMode: "html",
	})
	if err != nil {
		log.E("Couldn't send notification!")
		log.E(err.Error())
	}
}
