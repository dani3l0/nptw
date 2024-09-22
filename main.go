package main

import (
	"encoding/json"
	"fmt"
	"math"
	"nptw/config"
	"nptw/config/globals"
	"nptw/config/initialization"
	"nptw/providers/dlive"
	"nptw/providers/ia"
	"nptw/providers/telegram"
	"nptw/tools/ffmpeg"
	"nptw/tools/ytdlp"
	"nptw/utils/log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func main() {
	// Initialize project
	initialization.Init()

	// Some runtime variables
	wasStreaming := false
	readyToSend := false
	downloadRetries := 0
	replaysCache := ""

	// Main loop
	for {
		isLive, title, thumbnail, _ := dlive.GetStreamInfo()

		if isLive && !wasStreaming && config.Get().NotificationsEnabled {
			// Send notification to Telegram
			log.V("Ready to send notifications: " + strconv.FormatBool(readyToSend))
			text := fmt.Sprintf(
				"🇵🇱 Rozpoczął się Żywiec! 🇵🇱\n\n🐺 **%s** 🦎\n\n🔗 Link do DLive: https://dlive.tv/%s",
				strings.Replace(fmt.Sprintf("%v", title), "\n", "", -1), config.Get().DliveUsername,
			)
			if readyToSend {
				telegram.SendNotification(thumbnail, text)
			}
			wasStreaming = true

		} else if !isLive && wasStreaming && (config.Get().ReplaysEnabled || config.Get().IAEnabled) {
			log.I("Is not streaming now, but was streaming recently")
			log.I("Waiting for a while so DLive can properly archive the stream.")
			b2i := map[bool]int64{false: 1, true: 0}
			time.Sleep(time.Duration(5*b2i[config.Get().DebugMode]) * time.Minute)

			// Prepare filesystem
			log.I("It's time to grab the replay.")
			log.I("Creating cache path")
			os.MkdirAll(config.Get().CachePath, 0755)

			// Let's archive the stream
			replays, err := dlive.GetLastReplays()
			if err == nil && replays != replaysCache {
				resPath := "data.userByDisplayName.pastBroadcastsV2.list|0."

				// Grab information
				log.I("Grabbing information about archived stream")
				permlink := "https://dlive.tv/p/" + gjson.Get(replays, resPath+"permlink").Str
				title := gjson.Get(replays, resPath+"title").Str
				length := gjson.Get(replays, resPath+"length").Float()
				createdAt := gjson.Get(replays, resPath+"createdAt").Int()
				fullInfo, _ := ytdlp.GetInfo(permlink)
				var formats_raw map[string][]map[string]interface{}
				json.Unmarshal([]byte(fullInfo), &formats_raw)
				formats := formats_raw["formats"]
				targetSizeMB, targetFormat := .0, "none"
				log.I("Stream title: " + title)
				log.I(fmt.Sprint("Stream length: ", length, " seconds"))

				// Find proper format that meets our MaxReplaySizeMb requirement
				log.I("Finding proper format respecting our `MaxReplaySizeMb` limit")
				for _, f := range formats {
					tbr := f["tbr"].(float64)
					estimatedSizeMB := (tbr * length) / (8 * 1024)
					log.I(fmt.Sprint("Format `", f["format_id"], "`: bitrate ", tbr, " kbit/s, estimated size is ~", estimatedSizeMB, " MB"))
					if int(estimatedSizeMB) <= config.Get().MaxReplaySizeMb && estimatedSizeMB > targetSizeMB {
						targetSizeMB = estimatedSizeMB
						targetFormat = f["format_id"].(string)
					}
				}
				log.I("Finally, selected format `" + targetFormat + "`")
				if targetFormat == "none" && downloadRetries <= 3 {
					log.W("No video format selected. Trying again soon.")
					downloadRetries += 1
					time.Sleep(time.Minute * time.Duration(5*b2i[config.Get().DebugMode]))
					continue
				}

				// Download+Upload if small enough
				uploadSuccessful := false
				if targetSizeMB > 0 && targetFormat != "none" {
					// downloaded := ytdlp.Download(permlink, targetFormat)
					downloaded := false
					if downloaded {
						// Prepare message
						log.I("Prepare message")
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
							if !config.Get().ReplaysEnabled {
								ss := path.Join(config.Get().CachePath, globals.ScreenshotFilename)
								ffmpeg.Thumbnail(path.Join(config.Get().CachePath, globals.VideoFilename), int(length/4), ss)
							}
							ok, iaLink = ia.UploadReplay(title)
							if ok {
								iaLink = fmt.Sprintf("[Link do Archive.org](%s)", iaLink)
							} else {
								iaLink = "__Archive.org: błąd przesyłania__"
							}
						}

						// Upload to Telegram
						message := fmt.Sprintf(
							"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do DLive](%s)\n🥡 %s",
							title, duration, dayPol, formattedTime, permlink, iaLink,
						)
						uploadSuccessful = telegram.SendReplay(message, config.Get().ReplaysEnabled)

						// Cleanup
						if !config.Get().DebugMode {
							os.RemoveAll(config.Get().CachePath)
						}

					}
					if uploadSuccessful || downloadRetries > 3 {
						downloadRetries = 0
						wasStreaming = false
						replaysCache = replays
					} else {
						downloadRetries += 1
						log.W("Upload unsuccessful. Retry #" + strconv.Itoa(downloadRetries))
						continue
					}
				}
			}
		}

		if config.Get().DebugMode {
			time.Sleep(time.Second)
			os.Exit(0)
		}

		time.Sleep(time.Minute * time.Duration(config.Get().PollTime))
		readyToSend = true
	}
}
