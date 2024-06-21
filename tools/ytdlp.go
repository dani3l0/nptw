package tools

import (
	"nptw/config"
	"os/exec"
)

func YtDlpGetInfo() (string, error) {
	cmd := exec.Command("./bin/yt-dlp", "-J", "https://dlive.tv/"+config.Get().Username)
	stdout, err := cmd.Output()
	return string(stdout), err
}
