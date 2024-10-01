package dlive

import (
	"bytes"
	"io"
	"net/http"
	"nptw/config"
	"nptw/utils/log"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetLastReplay() (string, string, int64, int64, error) {
	// Weird graphigo payload
	jsonPath := "data.userByDisplayName.pastBroadcastsV2.list|0."
	jsonData := []byte(`{"query":"query($displayname: String!) { userByDisplayName(displayname: $displayname) { pastBroadcastsV2 { list { createdAt length permlink thumbnailUrl title } } } }","variables":{"displayname":"` + config.Get().DliveUsername + `"}}`)

	// Build a weird request
	log.V("Building LastReplay check request")
	req, err := http.NewRequest("POST", "https://graphigo.prd.dlive.tv/", bytes.NewBuffer(jsonData))
	if err != nil {
		log.E("Problem with building LastReplay request to graphigo: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// Do that weird request
	log.V("Doing LastReplay check request")
	resp, err := client.Do(req)
	if err != nil {
		log.E("Problem with doing LastReplay request to graphigo: ", err.Error())
	}

	// Get & parse result
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	json := string(data)
	log.V("LastReplay output:")
	log.V(json)

	// Gather information
	permlink := "https://dlive.tv/p/" + gjson.Get(json, jsonPath+"permlink").String()
	title := gjson.Get(json, jsonPath+"title").String()
	length := gjson.Get(json, jsonPath+"length").Int()
	createdAt := gjson.Get(json, jsonPath+"createdAt").Int()

	// Log stuff
	log.I("---------- Last replay info ----------")
	log.I("permlink:           ", permlink)
	log.I("title:              ", title)
	log.I("length in seconds:  ", strconv.Itoa(int(length)))
	log.I("creation date:      ", strconv.Itoa(int(createdAt)))

	return permlink, title, length, createdAt, err
}
