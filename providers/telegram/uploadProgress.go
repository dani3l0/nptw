package telegram

import (
	"fmt"
	"math"
	"nptw/config"
	"nptw/providers/rumble"
	"nptw/utils/log"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

// Structure for upload progress
type UploadProgress struct {
	Title           string
	Stage           int
	DownloadedBytes float64
	DownloadSize    int
	UploadedBytes   float64
	UploadSize      int
	Duration        int
	StartedAt       int64
	Error           string
	Done            bool
}

var ProgressUp UploadProgress
var stages = []string{"📦 Czekam na archiwizację", "💾 Pobieram i transkoduję", "📢 Wrzucam powtórkę na kanał", "✅ Powtórka wrzucona!"}

// Sends message with video upload progress
func SendUploadingProgress(entry rumble.Entry) {
	log.V("Sending message with progress")
	ProgressUp = UploadProgress{}
	ProgressUp.Duration = entry.Duration
	ProgressUp.Title = entry.Title
	ProgressUp.StartedAt = entry.Timestamp
	ProgressUp.Done = false
	sendOptions := &tg.MediaOptions{
		ParseMode: tg.HTML,
		Silent:    true,
		Caption:   uploadText(),
		Spoiler:   true,
	}
	msg, err := client.SendMedia(config.Get().ReplaysChannelId, entry.Thumbnail, sendOptions)
	if err != nil {
		log.E("Couldn't send progress message: ", err.Error())
		return
	}
	durr := time.Second * time.Duration(config.Get().DownloadProgressRefreshSec)
	for {
		time.Sleep(durr)
		msg.Edit(uploadText(), &tg.SendOptions{
			ParseMode: sendOptions.ParseMode,
		})
		if ProgressUp.Stage == len(stages)-1 {
			time.Sleep(durr)
			_, err = msg.Delete()
			if err != nil {
				log.E("Couldn't delete progress message: ", err.Error())
			}
			break
		}
	}
}

// Generates progress text
func uploadText() string {
	text := stages[ProgressUp.Stage]
	if ProgressUp.Error != "" {
		text = ProgressUp.Error
	}

	formattedTime := time.Unix(ProgressUp.StartedAt, 0).Format("02.01.2006 15:04")

	progress := ""
	switch ProgressUp.Stage {
	case 1:
		progress = progressLine(ProgressUp.DownloadedBytes, ProgressUp.DownloadSize)
	case 2:
		progress = progressLine(ProgressUp.UploadedBytes, ProgressUp.UploadSize)
	}

	return strings.ReplaceAll(strings.ReplaceAll(`
		<b>🐺 `+ProgressUp.Title+` 🦎</b>\n\n
		🕓 Czas trwania: <code>`+FormatDuration(ProgressUp.Duration)+`</code>\n
		🎥 Rozpoczęto: <code>`+formattedTime+`</code>\n\n
		<i>`+text+`</i>`+progress+`
	`, "\n", ""), "\\n", "\n")
}

// Formats time
func FormatDuration(length int) string {
	h := math.Floor(float64(length) / 3600)
	m := math.Floor(float64(length)/60) - h*60
	s := length % 60
	duration := fmt.Sprintf("%d:%02d:%02d", int(h), int(m), s)
	return duration
}

// Nice progress line during upload
func progressLine(current float64, target int) string {
	max := float64(target)
	pp := 100 * current / max
	line := ""
	for i := 0; i < 50; i++ {
		if int(math.Round(pp/2)) > i {
			line += "="
		} else {
			line += " "
		}
	}
	line = strings.Replace(line, " ", ">", 1)
	if pp > 100 {
		pp = 100
	}
	if current > max {
		current = max
	}
	return fmt.Sprintf(`\n\n<code> %.2f%%   %.0f MB / %.0f MB<code>\n<code>[%s]</code>`, pp, current, max, line)
}
