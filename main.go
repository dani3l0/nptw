package main

import (
	"nptw/config"
	"nptw/config/initialization"
	"nptw/providers/dlive"
	"nptw/providers/telegram"
	"nptw/utils/functions"
	"nptw/utils/log"
	"strconv"
	"time"
)

func main() {
	// Initialize project
	initialization.Init()

	// Some runtime variables
	wasLive := false
	readyToSend := false

	// Main loop
	for {
		isLive, title, _ := dlive.GetStreamInfo()

		if isLive && !wasLive && config.Get().NotificationsEnabled {
			// If live, send a notification to Telegram
			log.V("Ready to send notifications: " + strconv.FormatBool(readyToSend))
			thumbnail, _ := dlive.DlpInfo()
			if readyToSend {
				telegram.SendNotification(thumbnail, title)
			}
			wasLive = true

		} else if !isLive && wasLive && config.Get().ReplaysEnabled {
			// Wait for a moment before downloading archived stream
			log.I("Is not streaming now, but was streaming recently")
			log.I("Waiting for 10 minutes so DLive can properly archive the stream.")
			time.Sleep(10 * time.Minute)

			// Upload to Telegram|Archive.org
			functions.UploadReplay()
			wasLive = false
		}

		// Sleep for time specified in config
		time.Sleep(time.Minute * time.Duration(config.Get().PollTime))
		readyToSend = true
	}
}
