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

	debugN := config.Get().DebugNotifications
	debugR := config.Get().DebugReplays

	// Main loop
	for {
		isLive, title, _ := dlive.GetStreamInfo()

		// Debugging, development
		if debugN {
			isLive = true
			wasLive = false
			log.I("To debug notifications, target DLive channel MUST be streaming now!")
			if debugR {
				log.W("debug_replays flag is enabled, ignoring as notification and replay can't be sent at once")
			}
			readyToSend = true
		} else if debugR {
			isLive = false
			wasLive = true
		}

		if isLive && !wasLive && config.Get().NotificationsEnabled {
			// If live, send a notification to Telegram
			log.V("Ready to send notifications: " + strconv.FormatBool(readyToSend))
			var thumbnail string
			if readyToSend {
				// Try to get live thumbnail | max 5 times
				if config.Get().NotificationsLiveThumbnail {
					for i := 0; i < 5; i++ {
						thumbnail, _ = dlive.DlpInfo()
						if thumbnail == "" {
							log.W("No live thumbnail detected! Waiting for a minute (attempt ", strconv.Itoa(i+1), "/5)")
							time.Sleep(time.Minute)
						} else {
							break
						}
					}
				}
				telegram.SendNotification(thumbnail, title)
			}
			wasLive = true

		} else if !isLive && wasLive && config.Get().ReplaysEnabled {
			// Wait for a moment before downloading archived stream
			if !debugR {
				log.I("Is not streaming now, but was streaming recently")
				log.I("Waiting for 10 minutes so DLive can properly archive the stream.")
				time.Sleep(10 * time.Minute)
			} else {
				log.I("Uploading replay in debug mode, skipping wait time")
			}

			// Upload to Telegram|Archive.org
			functions.UploadReplay()
			wasLive = false
		}

		// Exit if debugging
		if debugN || debugR {
			return
		}

		// Sleep for time specified in config
		time.Sleep(time.Minute * time.Duration(config.Get().PollTime))
		readyToSend = true
	}
}
