package functions

import (
	"fmt"
	"io"
	"net/http"
	"nptw/config"
	"nptw/providers/telegram"
	"nptw/tools/ytdlp"
	"nptw/utils"
	"nptw/utils/log"
	"os"
	"path"
	"time"
)

// Uploads replay to Telegram
func UploadReplay(url string, title string, timestamp int64, length int, thumb string) {
	utils.PrepareCache()
	var t any
	tpath := path.Join(config.Get().CachePath, "thumbnail.png")
	if downloadPNG(thumb, tpath) == nil {
		t = tpath
	}
	downloaded := ytdlp.Download(url, length)

	if downloaded {
		// Prepare message
		log.I("Preparing message")
		startedAt := time.Unix(timestamp, 0)
		days := []string{"Niedziela", "Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota"}
		dayPol := days[startedAt.Weekday()]
		formattedTime := startedAt.Format("02.01.2006 15:04")

		// Upload to Telegram
		telegram.SendReplay(fmt.Sprintf(
			"🇵🇱 Żywiec - Powtórka 🇵🇱\n\n🐺 **%s** 🦎\n\n🕥 Czas trwania: `%s`\n📹 Rozpoczęto: `%s %s`\n\n📺 [Link do Rumbla](%s)",
			title, telegram.FormatDuration(length), dayPol, formattedTime, url,
		), t)

		// Cleanup files
		utils.CleanCache()
	} else {
		telegram.ProgressUp.Done = true
		telegram.ProgressUp.Stage = 3
		telegram.ProgressUp.Error = "Coś poszło nie tak..."
	}
}

func downloadPNG(url, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy data from response to file
	_, err = io.Copy(out, resp.Body)
	return err
}
