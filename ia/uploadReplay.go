package ia

import (
	"fmt"
	"net/url"
	"nptw/config"
	"nptw/utils"
	"os/exec"
	"path"
	"time"
)

func UploadReplay(title string) (bool, string) {
	Configure()
	currentDate := time.Now().Format("060102")
	filename := currentDate + " - " + path.Base(title) + ".mp4"
	utils.Log("Uploading video replay to archive.org")
	cmd := exec.Command(
		"./bin/ia", "upload",
		config.Get().IAFolderId, path.Join(config.Get().CachePath, config.Get().VideoFilename),
		"--metadata", "mediatype:movies",
		"-r", filename,
	)
	log, err := cmd.Output()
	if err == nil {
		utils.Log("Successfully uploaded video to archive.org")
		utils.Log(string(log))
	} else {
		utils.Err("Uploading video to archive.org failed!")
		utils.Err(err.Error())
		utils.Err(string(log))
	}
	return err == nil, fmt.Sprintf("https://archive.org/details/%s/%s", config.Get().IAFolderId, url.QueryEscape(filename))
}
