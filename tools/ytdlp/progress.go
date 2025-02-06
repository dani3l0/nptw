package ytdlp

import (
	"fmt"
	"nptw/config"
	"nptw/utils/log"
	"os"
	"strconv"
	"time"
)

func ProgressFunc(downloading *bool, filename string, expectedSize int) {
	var filesize_last float64 = 0
	progressRefreshTimeSec := config.Get().DownloadProgressRefreshSec
	log.I("Download progress: Refreshing each ", strconv.Itoa(progressRefreshTimeSec), " seconds")
	for *downloading {
		time.Sleep(time.Duration(progressRefreshTimeSec) * time.Second)
		f, err := os.Stat(filename)

		// Get downloaded file size
		var filesize float64
		if err != nil {
			filesize = 0
		} else {
			filesize = float64(f.Size()) / 1000 / 1000
		}
		speed := (filesize - filesize_last) * 1000 / float64(progressRefreshTimeSec)
		filesize_last = filesize

		// Log progress
		showProgress(filename, filesize, speed, expectedSize)
	}
}

func showProgress(filename string, loaded float64, speed float64, total int) {
	log.I("Downloading to '", filename, "': ",
		fmt.Sprint(float64(int(loaded*100))/100.0), " MB",
		" of approx. ",
		strconv.Itoa(total), " MB",
		" | ", fmt.Sprint(float64(int(speed*100))/100.0), "kB/s",
	)
}
