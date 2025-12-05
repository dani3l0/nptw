package rumble

// Structure for live information
type Entry struct {
	Id        string `json:"id"`
	Url       string `json:"original_url"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	IsLive    bool   `json:"is_live"`
	WasLive   bool   `json:"was_live"`
	Timestamp int64  `json:"release_timestamp,omitempty"`
}

// yt-dlp resp
type YtDlpResponse struct {
	Entries []Entry `json:"entries"`
}
