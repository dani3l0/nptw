package main

import (
	"nptw/tools"
	"nptw/utils"
)

func main() {
	utils.Check("Checking yt-dlp")
	if !tools.FfmpegCheck() || !tools.YtDlpCheck() {
		return
	}
}
