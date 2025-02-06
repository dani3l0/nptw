package config

// Default configuration
var config = Config{
	DliveUsername:                "nptvpl",
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
	LogLevel:                     2,
	DebugNotifications:           false,
	DebugReplays:                 false,
}
