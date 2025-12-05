package functions

import (
	"fmt"
	"nptw/providers/telegram"
	"nptw/tools/ytdlp"
	"nptw/utils"
	"nptw/utils/log"
	"time"
)

// Uploads replay to Telegram
func UploadReplay(url string, title string, timestamp int64, length int, thumb string) {
	utils.PrepareCache()
	downloaded := ytdlp.Download(url, length)

	if downloaded {
		// Prepare message
		log.I("Preparing message")
		startedAt := time.Unix(timestamp, 0)
		days := []string{"Niedziela", "Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota"}
		dayPol := days[startedAt.Weekday()]
		formattedTime := startedAt.Format("02.01.2006 15:04")

		// Upload to Telegram
		telegram.SendReplay(fmt.Sprintf(
			"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do Rumbla](%s)",
			title, telegram.FormatDuration(length), dayPol, formattedTime, url,
		), thumb)

		// Cleanup files
		utils.CleanCache()
	} else {
		telegram.ProgressUp.Done = true
		telegram.ProgressUp.Stage = 3
		telegram.ProgressUp.Error = "Coś poszło nie tak..."
	}
}
