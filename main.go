package main

import (
	"encoding/json"
	"fmt"
	"nptw/config"
	"nptw/telegram"
	"nptw/tools"
	"nptw/utils"
	"os"
	"strings"
	"time"
)

func main() {
	// Important checks
	os.MkdirAll("bin", 0750)
	if !config.Load() || !tools.FfmpegCheck() || !tools.YtDlpCheck() {
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

	// Main loop
	for {
		isStreaming := tools.IsStreaming()

		if isStreaming && !wasStreaming {
			fmt.Println("Started streaming!")

			// Get stream info
			var streamInfo map[string]interface{}
			streamInfoRaw, _ := tools.YtDlpGetInfo()
			json.Unmarshal([]byte(streamInfoRaw), &streamInfo)

			// Parse stream info
			isLiveForSure := streamInfo["is_live"]
			thumbnail := streamInfo["thumbnail"]
			title := streamInfo["fulltitle"]

			// Send notification to Telegram
			if isLiveForSure == true {
				text := fmt.Sprintf(
					"**🇵🇱 Rozpoczął się Żywiec! 🇵🇱**\n\n🐺 __%s__ 🦎\n\n➡️ Link do DLive: https://dlive.tv/%s",
					strings.Replace(fmt.Sprintf("%v", title), "\n", "", -1), config.Get().Username,
				)
				if readyToSend {
					telegram.SendNotification(thumbnail, text)
				}
				wasStreaming = true
			}

		} else if !isStreaming && wasStreaming {
			fmt.Println("Stream just ended")
			// yt-dlp download
			// telegram upload
			// archive.org upload
			// Purge
			if readyToSend {
				time.Sleep(time.Minute * 20)
			}
			wasStreaming = false
		}

		time.Sleep(time.Minute * 10)
		readyToSend = true
	}
}
