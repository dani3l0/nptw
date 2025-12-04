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
	FfmpegHwAccelDevice:          "/dev/dri/renderD128",
	FfmpegHwAccelType:            "cpu|qsv|vaapi|cuda",
	FfmpegThreads:                0,
	FfmpegHevc:                   false,
	PollTime:                     10,
	DownloadProgressRefreshSec:   5,
	LogLevel:                     strings.Join(log.LogLevels, "|"),
	DebugNotifications:           false,
	DebugReplays:                 false,
}
