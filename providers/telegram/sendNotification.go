package telegram

import (
	"fmt"
	"math/rand"
	"nptw/config"
	"nptw/utils/log"
	"path/filepath"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail string, title string, url string, dliveOk bool) {
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

	// Prepare dlive url
	dliveText := fmt.Sprintf(`<code>   </code><a href="%s">🥷 Dlive</a>`, config.Get().DliveUrl)
	if !dliveOk {
		dliveText = ""
	}

	// Prepare message
	text := fmt.Sprintf(`🇵🇱 Rozpoczął się Żywiec! 🇵🇱\n\n<b>🐺 %s 🦎</b>\n\n<a href="%s">📺 Rumble</a>%s`, title, url, dliveText)
	text = strings.ReplaceAll(text, "\\n", "\n")
	log.V(text)

	// Send message with preview photo
	_, err := client.SendMedia(config.Get().NotificationsChannelId, thumbnail, &tg.MediaOptions{
		Caption:   text,
		ParseMode: tg.HTML,
	})
	if err != nil {
		log.E("Couldn't send notification: ", err.Error())
	}
}
