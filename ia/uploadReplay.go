package ia

import (
	"fmt"
	"net/url"
	"nptw/config"
	"os/exec"
	"path"
)

func UploadReplay(file string) (bool, string) {
	Configure()
	cmd := exec.Command(
		"./bin/ia", "upload",
		config.Get().IAFolderId, file,
		"--metadata", "mediatype:movies",
	)
	res := cmd.Run() == nil
	return res, fmt.Sprintf("https://archive.org/details/%s/%s", config.Get().IAFolderId, url.PathEscape(path.Base(file)))
}
