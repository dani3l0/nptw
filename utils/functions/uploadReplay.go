package functions

import (
	"fmt"
	"math"
	"nptw/providers/telegram"
	"nptw/tools/ytdlp"
	"nptw/utils"
	"nptw/utils/log"
	"time"
)

// Uploads replay to Telegram
func UploadReplay(url string, title string, timestamp int64, length int) {
	utils.PrepareCache()
	downloaded := ytdlp.Download(url, length)

	if downloaded {
		// Prepare message
		log.I("Preparing message")
		h := math.Floor(float64(length) / 3600)
		m := math.Floor(float64(length)/60) - h*60
		s := length % 60
		duration := fmt.Sprintf("%d:%02d:%02d", int(h), int(m), s)
		startedAt := time.Unix(timestamp, 0)
		days := []string{"Niedziela", "Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota"}
		dayPol := days[startedAt.Weekday()]
		formattedTime := startedAt.Format("02.01.2006 15:04")

		// Upload to Telegram
		telegram.SendReplay(fmt.Sprintf(
			"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do DLive](%s)",
			title, duration, dayPol, formattedTime, url,
		))

		// Cleanup files
		utils.CleanCache()
	}
}
