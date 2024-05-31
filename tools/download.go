package tools

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filepath string) bool {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return false
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return false
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	return err != nil
}
