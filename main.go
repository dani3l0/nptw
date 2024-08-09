package main

import (
	"encoding/json"
	"fmt"
	"math"
	"nptw/config"
	"nptw/ia"
	"nptw/telegram"
	"nptw/tools"
	"nptw/tools/ffmpeg"
	"nptw/tools/ytdlp"
	"nptw/utils"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func main() {
	// Important checks
	os.MkdirAll("bin", 0750)
	if !config.Load() || !ytdlp.Check() || !ia.Check() {
		utils.Log("Oops... Something is wrong with your installation.")
		utils.Log("Try:")
		utils.Log("- removing 'bin' directory")
		utils.Log("- removing 'config.yaml' file")
		utils.Log("- setting proper permissions to your working dir")
		return
	}

	// Boot up
	utils.Log("Dependencies & configs are OK")
	telegram.Init()
	wasStreaming := false
	readyToSend := false
	replaysCache := ""
	utils.Log("Bot ready and running")

	// Main loop
	for {
		isStreaming := tools.IsStreaming()

		if config.Get().DebugMode {
			isStreaming = false
			wasStreaming = true
		}

		if isStreaming && !wasStreaming && config.Get().EnableNotifications {
			utils.Log("Stream started?")

			// Get stream info
			var streamInfo map[string]interface{}
			streamInfoRaw, _ := ytdlp.GetInfo("https://dlive.tv/" + config.Get().Username)
			utils.Log("Parsing stream information")
			json.Unmarshal([]byte(streamInfoRaw), &streamInfo)
			isLiveForSure := streamInfo["is_live"]
			thumbnail := streamInfo["thumbnail"]
			title := streamInfo["fulltitle"]

			// Send notification to Telegram
			if isLiveForSure == true {
				utils.Log("Is streaming now, for sure.")
				utils.Log("Ready to send notifications: " + strconv.FormatBool(readyToSend))
				text := fmt.Sprintf(
					"🇵🇱 Rozpoczął się Żywiec! 🇵🇱\n\n🐺 **%s** 🦎\n\n🔗 Link do DLive: https://dlive.tv/%s",
					strings.Replace(fmt.Sprintf("%v", title), "\n", "", -1), config.Get().Username,
				)
				if readyToSend {
					telegram.SendNotification(thumbnail, text)
				}
				wasStreaming = true
			}

		} else if !isStreaming && wasStreaming {
			utils.Log("Is not streaming now, but was streaming recently")
			utils.Log("Waiting for a while so DLive can properly archive the stream.")
			b2i := map[bool]int64{false: 1, true: 0}
			time.Sleep(time.Duration(15*b2i[config.Get().DebugMode]) * time.Minute)

			// Prepare filesystem
			utils.Log("It's time to grab the replay.")
			utils.Log("Creating cache path")
			os.MkdirAll(config.Get().CachePath, 0755)
			utils.Log("Creating symlink `replay.mp4`")
			ln_s := exec.Command("ln", "-s", path.Join(config.Get().CachePath, "replay.mp4"), "./replay.mp4")
			ln_s.Run()

			// Let's archive the stream
			replays, err := tools.GetLastReplays(1)
			if err == nil && replays != replaysCache {
				resPath := "data.userByDisplayName.pastBroadcastsV2.list|0."

				// Grab information
				utils.Log("Grabbing information about archived stream")
				permlink := "https://dlive.tv/p/" + gjson.Get(replays, resPath+"permlink").Str
				title := gjson.Get(replays, resPath+"title").Str
				length := gjson.Get(replays, resPath+"length").Float()
				createdAt := gjson.Get(replays, resPath+"createdAt").Int()
				fullInfo, _ := ytdlp.GetInfo(permlink)
				var formats_raw map[string][]map[string]interface{}
				json.Unmarshal([]byte(fullInfo), &formats_raw)
				formats := formats_raw["formats"]
				targetSizeMB, targetFormat := .0, "none"
				utils.Log("Stream title:    `" + title + "`")
				utils.Log(fmt.Sprint("Stream length:   ", length, " seconds"))

				// Find proper format that meets our MaxReplaySizeMb requirement
				utils.Log("Finding proper format respecting our `MaxReplaySizeMb` limit")
				for _, f := range formats {
					tbr := f["tbr"].(float64)
					estimatedSizeMB := (tbr * length) / (8 * 1024)
					utils.Log(fmt.Sprint("Format `", f["format_id"], "`: bitrate ", tbr, " kbit/s, estimated size is ~", estimatedSizeMB, " MB"))
					if int(estimatedSizeMB) <= config.Get().MaxReplaySizeMb && estimatedSizeMB > targetSizeMB {
						targetSizeMB = estimatedSizeMB
						targetFormat = f["format_id"].(string)
					}
				}
				utils.Log("Finally, selected format `" + targetFormat + "`")

				// Download+Upload if small enough
				if targetSizeMB > 0 && targetFormat != "none" {
					downloaded := ytdlp.Download(permlink, targetFormat)
					if downloaded {
						// Prepare message
						utils.Log("Prepare message")
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

							// Video filename for archive.org
							utils.Log("Symlinking a nice video name for archive.org")
							newname := path.Join(config.Get().CachePath, title+".mp4")
							os.Symlink(path.Join(config.Get().CachePath, "replay.mp4"), newname)

							// Generate thumbnail
							if !config.Get().EnableReplays {
								ss := path.Join(config.Get().CachePath, "screenshot.jpg")
								ffmpeg.Thumbnail(newname, int(length/4), ss)
							}

							// Upload stream
							ok, iaLink = ia.UploadReplay(newname)
							if ok {
								iaLink = fmt.Sprintf("[Link do Archive.org](%s)", iaLink)
							} else {
								iaLink = "__Archive.org: błąd przesyłania__"
							}
						}

						// Upload to Telegram
						if config.Get().EnableReplays {
							message := fmt.Sprintf(
								"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do DLive](%s)\n🥡 %s",
								title, duration, dayPol, formattedTime, permlink, iaLink,
							)
							telegram.SendReplay(title, message)
						}

						// Cleanup
						if !config.Get().DebugMode {
							os.RemoveAll(config.Get().CachePath)
						}

					}
					wasStreaming = false
					replaysCache = replays
				}
			}
		}

		time.Sleep(time.Minute * 15)
		readyToSend = true
	}
}
