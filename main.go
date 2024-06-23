package main

import (
	"encoding/json"
	"fmt"
	"math"
	"nptw/config"
	"nptw/telegram"
	"nptw/tools"
	"nptw/tools/ffmpeg"
	"nptw/tools/ytdlp"
	"nptw/utils"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func main() {
	// Important checks
	os.MkdirAll("bin", 0750)
	if !config.Load() || !ffmpeg.Check() || !ytdlp.Check() {
		utils.Log("Oops... Something is wrong with your installation.")
		utils.Log("Try:")
		utils.Log("- removing 'bin' directory")
		utils.Log("- removing 'config.yaml' file")
		utils.Log("- setting proper permissions to your working dir")
		return
	}

	// Boot up
	utils.Log("SUCCESS")
	telegram.Init()
	wasStreaming := false
	readyToSend := false
	replaysCache := ""
	utils.Log("Bot ready and running")

	// Main loop
	for {
		isStreaming := tools.IsStreaming()

		// Send notification if stream detected
		// Or download replay video if stream is over
		if isStreaming && !wasStreaming {
			fmt.Println("Started streaming!")

			// Get stream info
			var streamInfo map[string]interface{}
			streamInfoRaw, _ := ytdlp.GetInfo("https://dlive.tv/" + config.Get().Username)
			json.Unmarshal([]byte(streamInfoRaw), &streamInfo)

			// Parse stream info
			isLiveForSure := streamInfo["is_live"]
			thumbnail := streamInfo["thumbnail"]
			title := streamInfo["fulltitle"]

			// Send notification to Telegram
			if isLiveForSure == true {
				text := fmt.Sprintf(
					"🇵🇱 Rozpoczął się Żywiec! 🇵🇱\n\n🐺 **%s** 🦎\n\n🔗 Link do DLive: https://dlive.tv/%s",
					strings.Replace(fmt.Sprintf("%v", title), "\n", "", -1), config.Get().Username,
				)
				if readyToSend {
					telegram.SendNotification(thumbnail, text)
					time.Sleep(time.Minute * 15)
				}
				wasStreaming = true
			}

		} else if !isStreaming && wasStreaming {
			fmt.Println("Stream just ended")
			// yt-dlp download
			// telegram upload
			// archive.org upload
			// Purge

			replays, err := tools.GetLastReplays(1)
			if err == nil && replays != replaysCache {
				resPath := "data.userByDisplayName.pastBroadcastsV2.list|0."

				// Grab information
				permlink := "https://dlive.tv/p/" + gjson.Get(replays, resPath+"permlink").Str
				title := gjson.Get(replays, resPath+"title").Str
				length := gjson.Get(replays, resPath+"length").Float()
				createdAt := gjson.Get(replays, resPath+"createdAt").Int()
				fullInfo, _ := ytdlp.GetInfo(permlink)
				var formats_raw map[string][]map[string]interface{}
				json.Unmarshal([]byte(fullInfo), &formats_raw)
				formats := formats_raw["formats"]
				targetSizeMB, targetFormat := .0, "none"

				// Find proper format that can be sent to Telegram (less than 2GB file)
				for _, f := range formats {
					tbr := f["tbr"].(float64)
					estimatedSizeMB := (tbr * length) / (8 * 1024)
					if estimatedSizeMB <= 1536 && estimatedSizeMB > targetSizeMB {
						targetSizeMB = estimatedSizeMB
						targetFormat = f["format_id"].(string)
					}

				}

				if targetSizeMB > 0 && targetFormat != "none" {
					downloaded := ytdlp.Download(permlink, targetFormat)

					if downloaded {
						// Prepare message
						h := math.Floor(length / 3600)
						m := math.Floor(length/60) - h*60
						s := int(length) % 60
						duration := fmt.Sprintf("%d:%02d:%02d", int(h), int(m), s)
						startedAt := time.Unix(createdAt/1000, 0)
						days := []string{"Niedziela", "Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota"}
						dayPol := days[startedAt.Weekday()]
						formattedTime := startedAt.Format("02.01.2006  15:04")
						message := fmt.Sprintf("🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n📹 `%s`\n📅 `%s %s`\n\n🔗 Link do DLive: %s", title, duration, dayPol, formattedTime, permlink)

						fmt.Println(message)

						replaysCache = replays
						wasStreaming = false
					}
				}
			}
		}

		time.Sleep(time.Minute * 15)
		readyToSend = true
	}
}
