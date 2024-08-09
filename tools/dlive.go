package tools

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"nptw/config"
	"nptw/tools/ytdlp"
	"nptw/utils"
)

func IsStreaming() bool {
	utils.Log("Is streaming now? Checking")
	_, ok := ytdlp.GetInfo("https://dlive.tv/" + config.Get().Username)
	if ok == nil {
		utils.Log(config.Get().Username + " might be streaming now")
	} else {
		utils.Log(config.Get().Username + " seems to be offline")
	}
	return ok == nil
}

func GetLastReplays() (string, error) {
	utils.Log("Getting last replays")
	jsonPayload := fmt.Sprintf(
		`{"operationName":"LivestreamProfileReplay","variables":{"displayname":"%s","first":5},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"0417bdb00437901eec35ca7bf3e91ac5922f7ed2c5f5359e73a082102f71e810"}}}`,
		config.Get().Username,
	)

	// Build request
	request, _ := http.NewRequest("POST", "https://graphigo.prd.dlive.tv/", bytes.NewBufferString(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		utils.Warn("Sending request for last replays failed: " + err.Error())
		return "{}", err
	}

	// Read contents
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		utils.Warn("Failed decoding last replays data")
	}

	return string(data), err
}
