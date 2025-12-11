package main

import (
	"nptw/config"
	"nptw/config/initialization"
	"nptw/providers/dlive"
	"nptw/providers/rumble"
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
		liveData, err := rumble.GetInfo(config.Get().RumbleUrl)
		if err != nil {
			log.E("Wow, failed getting info. Waiting for a minute")
			time.Sleep(time.Minute)
			continue
		}
		streamData := liveData.Entries[0]

		// Debugging, development
		if debugN {
			streamData.IsLive = true
			wasLive = false
			log.I("To debug notifications, target channel MUST be streaming now!")
			if debugR {
				log.W("debug_replays flag is enabled, ignoring as notification and replay can't be sent at once")
			}
			readyToSend = true
		} else if debugR {
			streamData.IsLive = false
			wasLive = true
		}

		if streamData.IsLive && !wasLive && config.Get().NotificationsEnabled {
			// If live, send a notification to Telegram
			log.V("Ready to send notifications: " + strconv.FormatBool(readyToSend))
			var thumbnail string
			if readyToSend {
				// Try to get live thumbnail | max 5 times
				if config.Get().NotificationsLiveThumbnail {
					thumbnail = streamData.Thumbnail
				}
				telegram.SendNotification(thumbnail, streamData.Title, streamData.Url, dlive.IsLive())
			}
			wasLive = true

		} else if !streamData.IsLive && wasLive && config.Get().ReplaysEnabled {
			// Wait for a moment before downloading archived stream
			if !debugR {
				log.I("Is not streaming now, but was streaming recently")
				log.I("Waiting for 15 minutes so stream is properly archived")
				time.Sleep(15 * time.Minute)
			} else {
				log.I("Uploading replay in debug mode, skipping wait time")
			}

			// Get archived info
			sleep := time.Minute
			for range 60 {
				liveData, err := rumble.GetInfo(config.Get().RumbleUrl)
				if err != nil {
					log.E("Wow, failed getting info. Trying again in a moment. Error: ", err.Error())
					time.Sleep(sleep)
					continue
				}
				streamData = liveData.Entries[0]
				if streamData.Duration > 0 {
					break
				} else {
					log.E("Stream duration is 0s, waiting for appropriate data")
					time.Sleep(sleep)
				}
			}

			// Upload to Telegram
			if streamData.Duration > 0 {
				go telegram.SendUploadingProgress(streamData)
				if !debugR {
					time.Sleep(time.Minute)
				}
				functions.UploadReplay(streamData.Url, streamData.Title, streamData.Timestamp, streamData.Duration, streamData.Thumbnail)
			}
			wasLive = false
		}

		// Exit if debugging
		if debugN || debugR {
			if !debugN && debugR {
				time.Sleep(time.Minute)
			}
			return
		}

		// Sleep for time specified in config
		time.Sleep(time.Minute * time.Duration(config.Get().PollTime))
		readyToSend = true
	}
}
