package dlive

import (
	"bytes"
	"io"
	"net/http"
	"nptw/config"
	"nptw/utils/log"

	"github.com/tidwall/gjson"
)

func GetStreamInfo() (bool, string, error) {
	log.I("Checking if " + config.Get().DliveUsername + " is live ...")

	// Weird graphigo payload
	jsonPath := "data.userByDisplayName.livestream"
	jsonData := []byte(`{"query":"query($displayname: String!) { userByDisplayName(displayname: $displayname) { livestream { title } } }","variables":{"displayname":"` + config.Get().DliveUsername + `"}}`)

	// Build a weird request
	log.V("Building isLive check request")
	req, err := http.NewRequest("POST", "https://graphigo.prd.dlive.tv/", bytes.NewBuffer(jsonData))
	if err != nil {
		log.E("Problem with building isLive request to graphigo: ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// Do that weird request
	log.V("Doing isLive check request")
	resp, err := client.Do(req)
	if err != nil {
		log.E("Problem with doing isLive request to graphigo: ", err.Error())
	}

	// Get & parse result
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	json := string(data)
	log.V("isLive output:")
	log.V(json)

	// Find interesting variables
	isLive := gjson.Get(json, jsonPath).IsObject()
	title := "Offline"
	if isLive {
		title = gjson.Get(json, jsonPath+".title").String()
		log.I(config.Get().DliveUsername + " is live")
	} else {
		log.I(config.Get().DliveUsername + " is not live")
	}

	return isLive, title, err
}
