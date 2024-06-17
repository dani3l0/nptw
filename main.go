package main

import (
	"nptw/config"
	"nptw/telegram"
	"nptw/tools"
	"nptw/utils"
	"os"
)

func main() {
	// Important checks
	os.MkdirAll("bin", 0750)
	if !config.Load() || !tools.FfmpegCheck() || !tools.YtDlpCheck() {
		utils.Log("Oops... Something is wrong with your installation.")
		utils.Log("Try:")
		utils.Log("- removing 'bin' directory")
		utils.Log("- removing 'config.yaml' file")
		utils.Log("- setting proper permissions to your working dir")
		return
	}

	utils.Log("SUCCESS")

	telegram.Init()
}
