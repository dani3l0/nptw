package dlive

import (
	"bytes"
	"io"
	"net/http"
	"nptw/config"
	"nptw/utils/log"
	"strings"
)

func IsLive() bool {
	splt := strings.Split(config.Get().DliveUrl, "/")
	dliveUsername := splt[len(splt)-1]
	jsonData := []byte(`{"query":"query($displayname: String!) { userByDisplayName(displayname: $displayname) { livestream { title } } }","variables":{"displayname":"` + dliveUsername + `"}}`)

	// Build a weird request
	log.V("Doing isLive check request")
	req, _ := http.NewRequest("POST", "https://graphigo.prd.dlive.tv/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.E("Problem with doing isLive request to graphigo: ", err.Error())
		return false
	}

	// Get request result
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.E("Problem reading isLive data: ", err.Error())
		return false
	}

	// Simplify the response and just check whether online
	minijson := strings.ToLower(string(data))
	minijson = strings.ReplaceAll(minijson, "\n", "")
	minijson = strings.ReplaceAll(minijson, "\t", "")
	minijson = strings.ReplaceAll(minijson, " ", "")
	splot := strings.Split(minijson, `"livestream":null`)
	return len(splot) == 1
}
