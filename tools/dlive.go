package tools

func IsStreaming() bool {
	_, ok := YtDlpGetInfo()
	return ok == nil
}
