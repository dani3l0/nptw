package rumble

import (
	"encoding/json"
	"errors"
	"fmt"
	"nptw/utils/log"
	"os/exec"
	"time"
)

// Empty list err -___-
var ErrPlaylistEmpty = errors.New("playlist empty")

// Gets JSON-ed information about channel and last two replays
func GetInfo(url string) (YtDlpResponse, error) {
	var obj YtDlpResponse
	var err error
	for range 3 {
		c := exec.Command("./bin/yt-dlp", "-J", "-I", "1", url)
		output, errA := c.CombinedOutput()
		errB := json.Unmarshal(output, &obj)
		err = errors.Join(errA, errB)
		if err != nil {
			log.E("Couldn't get channel info: ", err.Error())
			log.E(string(output))
		}
		log.V(fmt.Sprintf("%+v", obj))
		if len(obj.Entries) == 0 {
			err = ErrPlaylistEmpty
		}
		if err == nil {
			break
		} else {
			log.E("rumble.GetInfo error: ", err.Error())
			time.Sleep(time.Second * 15)
		}
	}
	return obj, err
}
