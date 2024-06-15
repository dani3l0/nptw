package main

import (
	"nptw/config"
	"nptw/tools"
)

func main() {
	// Important checks
	if !tools.FfmpegCheck() || !tools.YtDlpCheck() || !config.Load() {
		return
	}
}
