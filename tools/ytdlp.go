package tools

import (
	"nptw/config"
	"os/exec"
)

func YtDlpGetInfo() (result string, err error) {
	cmd := exec.Command("./bin/yt-dlp", "-J", "https://dlive.tv/"+config.Get().Username)
	stdout, err := cmd.Output()
	return string(stdout), err
}
