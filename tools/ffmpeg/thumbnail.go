package ffmpeg

import (
	"nptw/utils"
	"os/exec"
	"strconv"
)

func Thumbnail(video string, target_sec int, output string) bool {
	utils.Log("Generating thumbnail for " + video + " at " + strconv.Itoa(target_sec) + " second")
	cmd := exec.Command("./bin/ffmpeg", "-y",
		"-ss", strconv.Itoa(target_sec),
		"-i", video,
		"-frames:v", "1", output,
	)

	ok := cmd.Run() == nil
	if ok {
		utils.Log("Thumbnail " + output + " generated successfully")
	} else {
		utils.Err("Thumbnail " + output + " generation failed")
	}
	return ok
}
