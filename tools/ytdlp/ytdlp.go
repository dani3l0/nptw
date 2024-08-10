package ytdlp

import (
	"nptw/config"
	"nptw/utils"
	"os"
	"os/exec"
	"path"
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
		"-f", format,
		"--ffmpeg-location", "bin",
		"--restrict-filenames",
		"-o", path.Join(config.Get().CachePath, config.Get().VideoFilename),
		"--cache-dir", path.Join(config.Get().CachePath, "cache"),
		url)

	log, _ := cmd.Output()
	stat, err := os.Stat(path.Join(config.Get().CachePath, config.Get().VideoFilename))
	ok := err == nil && stat.Size() > 0
	if ok {
		utils.Log("Downloading stream ended successfully")
		utils.Log(string(log))
	} else {
		utils.Err("Failed to download stream file")
		utils.Err(err.Error())
		utils.Err(string(log))
	}

	return ok
}
