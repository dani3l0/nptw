package rumble

import (
	"encoding/json"
	"errors"
	"fmt"
	"nptw/config"
	"nptw/utils/log"
	"os/exec"
)

// Empty list err -___-
var ErrPlaylistEmpty = errors.New("playlist empty")

// Gets JSON-ed information about channel and last two replays
func GetInfo(url string) (YtDlpResponse, error) {
	var obj YtDlpResponse
	playlistId := "1"
	if config.Get().DebugReplays {
		playlistId = "2"
	}
	c := exec.Command("./bin/yt-dlp", "-J", "-I", playlistId, url)
	output, errA := c.CombinedOutput()
	errB := json.Unmarshal(output, &obj)
	err := errors.Join(errA, errB)
	if err != nil {
		log.E("Couldn't get channel info: ", err.Error())
		log.E(string(output))
	}
	log.V(fmt.Sprintf("%+v", obj))
	if len(obj.Entries) == 0 {
		err = ErrPlaylistEmpty
	}
	return obj, err
}
