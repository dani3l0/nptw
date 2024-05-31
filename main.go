package main

import (
	"nptw/tools"
)

func main() {
	if !tools.FfmpegCheck() || !tools.YtDlpCheck() {
		return
	}
}
