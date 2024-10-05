package functions

import (
	"fmt"
	"math"
	"nptw/config"
	"nptw/providers/dlive"
	"nptw/providers/ia"
	"nptw/providers/telegram"
	"nptw/tools/ytdlp"
	"nptw/utils"
	"nptw/utils/log"
	"time"
)

func UploadReplay() { // Find information about last stream
	permlink, title, length, createdAt, err := dlive.GetLastReplay()

	if err == nil {
		// Download stream
		utils.PrepareCache()
		downloaded := ytdlp.Download(permlink, int(length))

		if downloaded {
			// Prepare message
			log.I("Preparing message")
			h := math.Floor(float64(length) / 3600)
			m := math.Floor(float64(length)/60) - h*60
			s := length % 60
			duration := fmt.Sprintf("%d:%02d:%02d", int(h), int(m), s)
			startedAt := time.Unix(createdAt/1000, 0)
			days := []string{"Niedziela", "Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota"}
			dayPol := days[startedAt.Weekday()]
			formattedTime := startedAt.Format("02.01.2006  15:04")

			// Upload to archive.org
			iaLink := ""
			if config.Get().IAEnabled {
				var ok bool

				// Taking hours to complete
				ok, iaLink = ia.UploadReplay(title)
				if ok {
					iaLink = fmt.Sprintf("\n🥡 [Link do Archive.org](%s)", iaLink)
				} else {
					iaLink = "\n🥡 __Archive.org: błąd przesyłania__"
				}
			}

			// Upload to Telegram
			telegram.SendReplay(fmt.Sprintf(
				"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do DLive](%s)%s",
				title, duration, dayPol, formattedTime, permlink, iaLink,
			))

			// Cleanup files
			utils.CleanCache()
		}
	}
}
