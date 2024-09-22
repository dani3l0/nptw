package dlive

import (
	"fmt"
	"nptw/utils/log"

	"github.com/tidwall/gjson"
)

func ParseLastReplay() (string, string, float64, int64, error) {
	var err error
	resPath := "data.userByDisplayName.pastBroadcastsV2.list|0."

	// Fetch data
	replays, err := GetLastReplays()
	if err != nil {
		log.E("Couldn't get last replays: ", err.Error())
	}
	log.V("GetLastReplays() output:")
	log.V(replays)

	// Grab & parse information
	log.I("Grabbing information about archived stream")
	permlink := "https://dlive.tv/p/" + gjson.Get(replays, resPath+"permlink").Str
	title := gjson.Get(replays, resPath+"title").Str
	length := gjson.Get(replays, resPath+"length").Float()
	createdAt := gjson.Get(replays, resPath+"createdAt").Int()
	log.I("Stream title: " + title)
	log.I(fmt.Sprint("Stream length: ", length, " seconds"))

	return permlink, title, length, createdAt, err
}
