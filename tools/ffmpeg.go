package tools

import (
	"nptw/utils"
	"os"
)

func FfmpegCheck() bool {
	utils.Check("Checking ffmpeg")
	info, err := os.Stat("./bin/ffmpeg")
	if os.IsExist(err) {
		return true
	}
	if !FfmpegGet() {
		return false
	}
	return info.IsDir()
}

func FfmpegGet() bool {
	return false
}
