package telegram

import (
	"fmt"
	"math/rand"
	"nptw/config"
	"nptw/utils/log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func SendNotification(thumbnail string, title string, url string, timestamp int64, dliveOk bool) {
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

	// Prepare how long ago started
	now := time.Now()
	timestampDifference := now.Unix() - timestamp

	// Prepare dlive url
	dliveText := fmt.Sprintf(`<code>   </code><a href="%s">🥷 Dlive</a>`, config.Get().DliveUrl)
	if dliveOk {
		dliveText = ""
	}

	// Prepare message
	text := fmt.Sprintf(`<b>🇵🇱 Rozpoczął się Żywiec! 🇵🇱</b>\n\n<i>🐺 %s 🦎</i>\n\n<a href="%s">📺 Rumble</a>%s\n\n<i>%s</i>`, title, url, dliveText, timePlural(int(timestampDifference)))
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

// Nice time string
func timePlural(seconds int) string {
	unit := "sekund"
	if seconds >= 60 {
		unit = "minut"
		seconds /= 60
	}
	if seconds >= 60 {
		unit = "godzin"
		seconds /= 60
	}
	r := seconds % 10
	if seconds == 1 || seconds == 0 {
		unit += "ę"
	} else if (r == 2 || r == 3 || r == 4) && !(10 < seconds && seconds < 20) {
		unit += "y"
	}
	if seconds > 1 {
		unit = strconv.Itoa(seconds) + " " + unit
	}
	return fmt.Sprintf("%s temu", unit)
}
