package ffmpeg

import (
	"os/exec"
	"strconv"
)

func Thumbnail(video string, target_sec int, output string) bool {
	cmd := exec.Command("./bin/ffmpeg", "-y",
		"-ss", strconv.Itoa(target_sec),
		"-i", video,
		"-frames:v", "1", output,
	)

	return cmd.Run() == nil
}
