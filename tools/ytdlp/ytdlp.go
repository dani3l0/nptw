package ytdlp

import (
	"nptw/config"
	"nptw/utils"
	"os"
	"os/exec"
	"path"
	"strings"
)

// Get info about current stream
func GetInfo(url string) (string, error) {
	utils.Log("yt-dlp is getting info about " + url)
	cmd := exec.Command("./bin/yt-dlp", "-J", url)
	stdout, err := cmd.Output()
	return string(stdout), err
}

// Download last stream, and save it to cache folder
func Download(url string, format string) bool {
	utils.Log("Downloading raw stream file from " + url + " with format " + format)

	os.MkdirAll(path.Join(config.Get().CachePath, "cache"), 0755)
	cmd := exec.Command(
		"./bin/yt-dlp",
		"--quiet",
		"-f", format,
		"--ffmpeg-location", "bin",
		"--restrict-filenames",
		"-o", config.Get().CachePath+"/replay.mp4",
		"--cache-dir", path.Join(config.Get().CachePath, "cache"),
		url)

	log, err := cmd.Output()
	if err == nil {
		utils.Log("Downloading stream ended successfully")
	} else {
		utils.Err("Failed to download stream file")
	}
	utils.Log("yt-dlp output:")
	for _, v := range strings.Split(string(log), "\n") {
		utils.Log(v)
	}

	return err == nil
}
