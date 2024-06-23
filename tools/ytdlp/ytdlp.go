package ytdlp

import (
	"nptw/config"
	"os"
	"os/exec"
	"path"
)

// Get info about current stream
func GetInfo(url string) (string, error) {
	cmd := exec.Command("./bin/yt-dlp", "-J", url)
	stdout, err := cmd.Output()
	return string(stdout), err
}

// Download last stream, and save it to cache folder
func Download(url string, format string) bool {
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

	_, err := cmd.Output()
	return err == nil
}
