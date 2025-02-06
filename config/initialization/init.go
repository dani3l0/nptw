package initialization

import (
	"nptw/config"
	"nptw/providers/telegram"
	"nptw/tools/ffmpeg"
	"nptw/tools/ytdlp"
	"nptw/utils/log"
	"os"
)

func Init() {
	// Check config & get dependencies
	os.MkdirAll("bin", 0750)
	if !config.Load() || !ffmpeg.Check() || !ytdlp.Check() {
		log.E("Oops... Something is wrong with your installation.")
		log.E("Try:")
		log.E("- removing 'bin' directory")
		log.E("- removing 'config.yaml' file")
		log.E("- setting proper permissions to your working dir")
		return
	}
	log.I("Dependencies & configs are OK")

	// Telegram connection
	telegram.Init()
}
