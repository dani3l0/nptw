package config

import (
	"nptw/config/globals"
	"nptw/utils/log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DliveUsername                string `yaml:"dlive_username"`
	TelegramBotToken             string `yaml:"telegram_bot_token"`
	TelegramApiId                int    `yaml:"telegram_api_id"`
	TelegramApiHash              string `yaml:"telegram_api_hash"`
	RespondToUserMessagesEnabled bool   `yaml:"respond_to_user_messages"`
	NotificationsEnabled         bool   `yaml:"notifications_enabled"`
	NotificationsChannelId       int    `yaml:"notifications_channel_id"`
	NotificationsLiveThumbnail   bool   `yaml:"notifications_live_thumb"`
	ReplaysEnabled               bool   `yaml:"replays_enabled"`
	ReplaysChannelId             int    `yaml:"replays_channel_id"`
	MaxReplaySizeMb              int    `yaml:"max_replay_size_mb"`
	CachePath                    string `yaml:"cache_path"`
	FfmpegHwAccelType            string `yaml:"ffmpeg_hwaccel_type"`
	FfmpegHwAccelDevice          string `yaml:"ffmpeg_hwaccel_device"`
	FfmpegThreads                int    `yaml:"ffmpeg_threads"`
	FfmpegHevc                   bool   `yaml:"ffmpeg_hevc"`
	PollTime                     int    `yaml:"poll_time"`
	DownloadProgressRefreshSec   int    `yaml:"download_progress_refresh_sec"`
	LogLevel                     int    `yaml:"log_level"`
	DebugNotifications           bool   `yaml:"debug_notifications"`
	DebugReplays                 bool   `yaml:"debug_replays"`
}

// Get config somewhere in code
func Get() Config {
	return config
}

// Open file and if exists, read configuration
func Load() bool {
	log.LogLevel = 10
	file, err := os.ReadFile(globals.ConfigFileName)

	// Attempt to create new configuration file
	if err != nil {
		log.W("Looks like it's the first run.")
		log.W("Creating new config file")

		// Write default config to filesystem
		d, _ := yaml.Marshal(&config)
		err := os.WriteFile(globals.ConfigFileName, []byte(d), 0640)
		if err != nil {
			log.E("Couldn't create new file. Make sure you have proper permissions to do so.")
			os.Exit(1)
		}

		log.I("Config file '" + globals.ConfigFileName + "' with default values created successfully.")
		log.W("Before proceeding, please adjust it to your likings.")
		log.W("Make sure to provide valid Telegram creds!")
		log.W("----- !!! Also, remember to send a message right after bot starts for the first time !!! -----")
		os.Exit(0)
	}

	// Try to read existing config
	log.I("Loading config file ...")
	config = Config{}
	err = yaml.Unmarshal([]byte(file), &config)
	if err != nil {
		log.E("Failed reading configuration file: ", err.Error())
	} else {
		log.I("Config file loaded successfully!")
	}
	log.LogLevel = config.LogLevel
	return err == nil
}
