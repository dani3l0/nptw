package ytdlp

import (
	"nptw/utils"
	"os"
	"os/exec"
)

// YtDlp check and auto-download
func Check() bool {
	exists := Exists()
	if !exists {
		Install()
		exists = Exists()
	}
	return exists
}

// Check if yt-dlp is available
func Exists() bool {
	utils.Check("Checking yt-dlp")

	info, err := os.Stat("./bin/yt-dlp")
	ok := !os.IsNotExist(err)
	if ok {
		ok = !info.IsDir()
	}

	utils.OkFail(ok)
	return ok
}

// Download yt-dlp locally
func Install() bool {
	utils.Check("Installing yt-dlp")
	cmd := exec.Command("bash", "-c", `
		wget -O ./bin/yt-dlp https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux
		chmod +x ./bin/yt-dlp
	`)

	ok := cmd.Run() == nil
	utils.OkFail(ok)

	return ok
}
