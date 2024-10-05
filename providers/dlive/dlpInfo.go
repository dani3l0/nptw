package dlive

import (
	"nptw/config"
	"nptw/utils/log"
	"os/exec"

	"github.com/tidwall/gjson"
)

func DlpInfo() (string, error) {
	log.I("Getting full info about stream via yt-dlp")
	json, err := exec.Command("./bin/yt-dlp", "-J", "https://dlive.tv/"+config.Get().DliveUsername).Output()
	parsed := gjson.ParseBytes(json)
	log.V("yt-dlp response:")
	log.V(parsed.Raw)
	thumbnail := gjson.Get(parsed.Raw, "thumbnails|0.url").String()
	return thumbnail, err
}
