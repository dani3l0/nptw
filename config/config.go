package config

import (
	"nptw/utils"
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
	ReplaysEnabled               bool   `yaml:"replays_enabled"`
	ReplaysChannelId             int    `yaml:"replays_channel_id"`
	IAEnabled                    bool   `yaml:"internet_archive_enabled"`
	IAEmail                      string `yaml:"internet_archive_email"`
	IAPassword                   string `yaml:"internet_archive_password"`
	IAFolderId                   string `yaml:"internet_archive_folder"`
	MaxReplaySizeMb              int    `yaml:"max_replay_size_mb"`
	CachePath                    string `yaml:"cache_path"`
	FfmpegHwAccelType            string `yaml:"ffmpeg_hwaccel_type"`
	FfmpegHwAccelDevice          string `yaml:"ffmpeg_hwaccel_device"`
	FfmpegReplayWidthPixels      int    `yaml:"ffmpeg_replay_width_pixels"`
	PollTime                     int    `yaml:"poll_time"`
	LogLevel                     int    `yaml:"log_level"`
	DebugMode                    bool   `yaml:"debug_mode"`
}

// Get config somewhere in code
func Get() Config {
	return config
}

// Open file and if exists, read configuration
func Load() bool {
	log.LogLevel = 10
	file, err := os.ReadFile(fileName)

	// Attempt to create new configuration file
	if err != nil {
		log.I("Looks like it's the first run.")
		utils.Check("Creating new config file")

		// Write default config to filesystem
		d, _ := yaml.Marshal(&config)
		err := os.WriteFile(fileName, []byte(d), 0640)
		if err != nil {
			utils.OkFail(false)
			log.E("Couldn't create new file. Make sure you have proper permissions to do so.")
			os.Exit(1)
		}

		utils.OkFail(true)
		log.W("Before proceeding, please adjust '" + fileName + "' file to your likings.")
		log.W("Make sure to provide valid Telegram creds!")
		log.W("Also, remember to send a message right after bot starts for the first time.")
		os.Exit(0)
	}

	// Try to read existing config
	utils.Check("Loading config file")
	config = Config{}
	err = yaml.Unmarshal([]byte(file), &config)
	ok := err == nil
	utils.OkFail(ok)
	log.LogLevel = config.LogLevel
	return ok
}
