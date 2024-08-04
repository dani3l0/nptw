package ia

import (
	"fmt"
	"net/url"
	"nptw/config"
	"nptw/utils"
	"os/exec"
	"path"
)

func UploadReplay(file string) (bool, string) {
	Configure()
	utils.Log("Uploading video replay to archive.org")
	cmd := exec.Command(
		"./bin/ia", "upload",
		config.Get().IAFolderId, file,
		"--metadata", "mediatype:movies",
	)
	res := cmd.Run() == nil
	if res {
		utils.Log("Successfully uploaded video to archive.org")
	} else {
		utils.Err("Uploading video to archive.org failed")
	}
	return res, fmt.Sprintf("https://archive.org/details/%s/%s", config.Get().IAFolderId, url.PathEscape(path.Base(file)))
}
