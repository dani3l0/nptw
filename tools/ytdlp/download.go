package ytdlp

import (
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"os/exec"
	"strconv"
	"strings"
)

func Download(url string, length int) bool {
	// Calculate appropriate bitrate
	maxSizeKb := config.Get().MaxReplaySizeMb * 1000
	maxKbPerSec := maxSizeKb / length
	maxKbitPerSec := maxKbPerSec * 8
	if maxKbitPerSec > globals.MaxBitrate {
		maxKbitPerSec = globals.MaxBitrate
		maxKbPerSec = maxKbitPerSec / 8
	}
	maxVideoKbitPerSec := maxKbitPerSec - globals.AudioBitrate

	// Log calculations
	log.V("Combined bitrate:  ", strconv.Itoa(maxKbitPerSec), " kbps")
	log.V("Video bitrate:     ", strconv.Itoa(maxVideoKbitPerSec), " kbps")
	log.V("Approx. filesize:  ", strconv.Itoa(maxKbPerSec*length/1000), " MB")

	// Let's transcode
	cmd := GetYtDlpFfmpegCmd(url, maxVideoKbitPerSec)
	log.I("Generated ffmpeg command:")
	log.I(strings.Join(cmd, " "))
	c := exec.Command("/bin/sh", "-c", strings.Join(cmd, " "))
	output, err := c.CombinedOutput()
	log.V(string(output))
	if err != nil {
		log.E("Transcoding problem: ", err.Error())
		log.E(string(output))
	}

	return err == nil
}
