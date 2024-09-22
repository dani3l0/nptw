package ia

import (
	"fmt"
	"net/url"
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"os/exec"
	"path"
	"time"
)

func UploadReplay(title string) (bool, string) {
	Configure()
	currentDate := time.Now().Format("060102")
	filename := currentDate + " - " + path.Base(title) + ".mp4"
	log.I("Uploading video replay to archive.org")
	cmd := exec.Command(
		"./bin/ia", "upload",
		config.Get().IAFolderId, path.Join(config.Get().CachePath, globals.VideoFilename),
		"--metadata", "mediatype:movies",
		"-r", filename,
	)
	output, err := cmd.Output()
	if err == nil {
		log.I("Successfully uploaded video to archive.org")
		log.I(string(output))
	} else {
		log.E("Uploading video to archive.org failed!")
		log.E(err.Error())
		log.E(string(output))
	}
	return err == nil, fmt.Sprintf("https://archive.org/details/%s/%s", config.Get().IAFolderId, url.QueryEscape(filename))
}
