package ytdlp

import (
	"nptw/config"
	"nptw/config/globals"
	"nptw/tools"
	"nptw/utils/log"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func Download(url string, length int) bool {
	// Calculate appropriate bitrate
	maxSizeKb := config.Get().MaxReplaySizeMb * 1000
	maxKbitPerSec := maxSizeKb / length * 8
	if maxKbitPerSec > globals.MaxBitrate {
		maxKbitPerSec = globals.MaxBitrate
	} else if maxKbitPerSec < globals.MinBitrate {
		maxKbitPerSec = globals.MinBitrate
	}
	maxKbPerSec := maxKbitPerSec / 8
	maxVideoKbitPerSec := maxKbitPerSec - globals.AudioBitrate
	sizeMb := maxKbPerSec * length / 1000

	// Log calculations
	log.I("Combined bitrate:  ", strconv.Itoa(maxKbitPerSec), " kbps")
	log.I("Video bitrate:     ", strconv.Itoa(maxVideoKbitPerSec), " kbps")
	log.I("Approx. filesize:  ", strconv.Itoa(sizeMb), " MB")

	if sizeMb > config.Get().MaxReplaySizeMb {
		log.E("Replay video too big! Expected size is ", strconv.Itoa(sizeMb), " MB while maximum allowed size for replay is ", strconv.Itoa(config.Get().MaxReplaySizeMb), " MB")
		return false
	}

	// Progress function
	downloading := true
	go ProgressFunc(&downloading, path.Join(config.Get().CachePath, globals.VideoFilename), maxKbPerSec*length/1000)

	// Let's transcode
	cmd := tools.GetYtDlpFfmpegCmd(url, maxVideoKbitPerSec)
	log.I("Generated ffmpeg command:")
	log.I(strings.Join(cmd, " "))
	c := exec.Command("/bin/sh", "-c", strings.Join(cmd, " "))
	output, err := c.CombinedOutput()
	downloading = false
	log.V(string(output))
	if err != nil {
		log.E("Transcoding problem: ", err.Error())
		log.E(string(output))
	}

	return err == nil
}
