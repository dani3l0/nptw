package config

import (
	"nptw/utils/log"
	"strings"
)

// Default configuration
var config = Config{
	RumbleUrl:                    "https://rumble.com/c/RodacyKamraciPL",
	DliveUrl:                     "https://dlive.tv/nptvpl",
	TelegramBotToken:             "ur_token_goes_here",
	TelegramApiId:                123456,
	TelegramApiHash:              "some_very_long_secret_hash",
	RespondToUserMessagesEnabled: true,
	NotificationsEnabled:         true,
	NotificationsChannelId:       123456789,
	NotificationsLiveThumbnail:   true,
	ReplaysEnabled:               true,
	ReplaysChannelId:             456789,
	MaxReplaySizeMb:              1600,
	CachePath:                    "./cache",
	YtDlpThreads:                 1,
	FfmpegHwAccelType:            "cpu|qsv|vaapi|cuda",
	FfmpegThreads:                0,
	FfmpegFormat:                 "h264|hevc|av1",
	PollTime:                     10,
	DownloadProgressRefreshSec:   10,
	LogLevel:                     strings.Join(log.LogLevels, "|"),
	DebugNotifications:           false,
	DebugReplays:                 false,
}
