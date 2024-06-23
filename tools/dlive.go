package tools

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"nptw/config"
	"nptw/tools/ytdlp"
)

func IsStreaming() bool {
	_, ok := ytdlp.GetInfo("https://dlive.tv/" + config.Get().Username)
	return ok == nil
}

func GetLastReplays(n int) (string, error) {
	jsonPayload := fmt.Sprintf(
		`{"operationName":"LivestreamProfileReplay","variables":{"displayname":"%s","first":%d},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"0417bdb00437901eec35ca7bf3e91ac5922f7ed2c5f5359e73a082102f71e810"}}}`,
		config.Get().Username, n,
	)

	// Build request
	request, _ := http.NewRequest("POST", "https://graphigo.prd.dlive.tv/", bytes.NewBufferString(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "{}", err
	}

	// Read contents
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	return string(data), err
}
