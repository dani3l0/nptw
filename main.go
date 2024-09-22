package main

import (
	"fmt"
	"math"
	"nptw/config"
	"nptw/config/globals"
	"nptw/config/initialization"
	"nptw/providers/dlive"
	"nptw/providers/ia"
	"nptw/providers/telegram"
	"nptw/tools/ffmpeg"
	"nptw/utils"
	"nptw/utils/log"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initialize project
	initialization.Init()

	// Some runtime variables
	wasStreaming := false
	readyToSend := false

	// Main loop
	for {
		isLive, title, thumbnail, _ := dlive.GetStreamInfo()

		if isLive && !wasStreaming && config.Get().NotificationsEnabled {
			// If live, send a notification to Telegram
			log.V("Ready to send notifications: " + strconv.FormatBool(readyToSend))
			if readyToSend {
				telegram.SendNotification(thumbnail, fmt.Sprintf(
					"🇵🇱 Rozpoczął się Żywiec! 🇵🇱\n\n🐺 **%s** 🦎\n\n🔗 Link do DLive: https://dlive.tv/%s",
					strings.Replace(fmt.Sprintf("%v", title), "\n", "", -1), config.Get().DliveUsername,
				))
			}
			wasStreaming = true

		} else if !isLive && wasStreaming && (config.Get().ReplaysEnabled || config.Get().IAEnabled) {
			// Wait for a moment before downloading archived stream
			log.I("Is not streaming now, but was streaming recently")
			log.I("Waiting for 5 minutes so DLive can properly archive the stream.")
			time.Sleep(time.Duration(5) * time.Minute)

			// Find information about last stream
			permlink, title, length, createdAt, err := dlive.ParseLastReplay()

			if err != nil {
				// Download stream
				utils.PrepareCache()
				downloaded := true
				if downloaded {
					// Prepare message
					log.I("Preparing message")
					h := math.Floor(length / 3600)
					m := math.Floor(length/60) - h*60
					s := int(length) % 60
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

		// Sleep for time specified in config
		time.Sleep(time.Minute * time.Duration(config.Get().PollTime))
		readyToSend = true
	}
}
