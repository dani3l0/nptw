package dlive

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"nptw/config"
	"nptw/utils/log"
)

func GetLastReplays() (string, error) {
	log.I("Getting last replays")
	jsonPayload := fmt.Sprintf(
		`{"operationName":"LivestreamProfileReplay","variables":{"displayname":"%s","first":5},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"0417bdb00437901eec35ca7bf3e91ac5922f7ed2c5f5359e73a082102f71e810"}}}`,
		config.Get().DliveUsername,
	)

	// Build request
	request, _ := http.NewRequest("POST", "https://graphigo.prd.dlive.tv/", bytes.NewBufferString(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.E("Sending request for last replays failed: " + err.Error())
		return "{}", err
	}

	// Read contents
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.E("Failed decoding last replays data")
	}

	return string(data), err
}
