package ytdlp

import (
	"nptw/utils/log"
	"os"
	"os/exec"
)

// yt-dlp check and auto-download
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
	info, err := os.Stat("./bin/yt-dlp")
	ok := !os.IsNotExist(err)
	if ok && !info.IsDir() {
		log.I("yt-dlp binary file seems to exist.")
	} else if !ok {
		log.W("yt-dlp binary file does not exist. Downloading now.")
	} else {
		log.E("Path for an yt-dlp binary file is broken.")
		log.E("Try removing whole `bin` directory, setting proper permissions or providing valid path for assets.")
	}
	return ok
}

// Download yt-dlp locally
func Install() bool {
	log.I("Installing yt-dlp")
	cmd := exec.Command("bash", "-c", `
		wget -O ./bin/yt-dlp https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux
		chmod +x ./bin/yt-dlp
	`)

	ok := cmd.Run() == nil

	if ok {
		log.I("yt-dlp installed successfully.")
	} else {
		log.E("Couldn't install yt-dlp properly.")
	}

	return ok
}
