package functions

import (
	"fmt"
	"math"
	"nptw/config"
	"nptw/config/globals"
	"nptw/providers/dlive"
	"nptw/providers/ia"
	"nptw/providers/telegram"
	"nptw/tools/ffmpeg"
	"nptw/utils"
	"nptw/utils/log"
	"path"
	"time"
)

func UploadReplay() { // Find information about last stream
	permlink, title, length, createdAt, err := dlive.GetLastReplay()

	if err != nil {
		// Download stream
		utils.PrepareCache()
		downloaded := true

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
			iaLink := "__Archive.org: funkcja wyłączona__"
			if config.Get().IAEnabled {
				var ok bool

				// If not uploading to Telegram, generate a nice thumbnail to be sent as photo
				if !config.Get().ReplaysEnabled {
					ss := path.Join(config.Get().CachePath, globals.ScreenshotFilename)
					ffmpeg.Thumbnail(path.Join(config.Get().CachePath, globals.VideoFilename), int(length/4), ss)
				}

				// Taking hours to complete
				ok, iaLink = ia.UploadReplay(title)
				if ok {
					iaLink = fmt.Sprintf("[Link do Archive.org](%s)", iaLink)
				} else {
					iaLink = "__Archive.org: błąd przesyłania__"
				}
			}

			// Upload to Telegram
			telegram.SendReplay(fmt.Sprintf(
				"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do DLive](%s)\n🥡 %s",
				title, duration, dayPol, formattedTime, permlink, iaLink,
			), config.Get().ReplaysEnabled)

			// Cleanup files
			utils.CleanCache()
		}
	}
}
