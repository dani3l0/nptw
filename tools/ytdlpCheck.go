package tools

import (
	"nptw/utils"
	"os"
	"os/exec"
)

// YtDlp check and auto-download
func YtDlpCheck() bool {
	exists := YtDlpExists()
	if !exists {
		YtDlpInstall()
		exists = YtDlpExists()
	}
	return exists
}

// Check if yt-dlp is available
func YtDlpExists() bool {
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
func YtDlpInstall() bool {
	utils.Check("Installing yt-dlp")
	cmd := exec.Command("bash", "-c", `
		wget -O ./bin/yt-dlp https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux
		chmod +x ./bin/yt-dlp
	`)

	ok := cmd.Run() == nil
	utils.OkFail(ok)

	return ok
}
