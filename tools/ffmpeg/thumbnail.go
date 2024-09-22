package ffmpeg

import (
	"nptw/utils/log"
	"os/exec"
	"strconv"
)

func Thumbnail(video string, target_sec int, photo string) bool {
	log.I("Generating thumbnail for " + video + " at " + strconv.Itoa(target_sec) + " second")
	cmd := exec.Command("./bin/ffmpeg", "-y",
		"-ss", strconv.Itoa(target_sec),
		"-i", video,
		"-frames:v", "1", photo,
	)

	output, err := cmd.Output()
	if err == nil {
		log.I("Thumbnail " + photo + " generated successfully")
		log.I(string(output[:]))
	} else {
		log.E("Thumbnail " + photo + " generation failed")
		log.E(err.Error())
		log.E(string(output[:]))
	}
	return err == nil
}
