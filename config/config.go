package config

import (
	"nptw/utils"
	"os"

	"gopkg.in/yaml.v3"
)

var fileName = "config.yaml"

type Config struct {
	TelegramBotToken    string `yaml:"telegram_bot_token"`
	TelegramApiId       int    `yaml:"telegram_api_id"`
	TelegramApiHash     string `yaml:"telegram_api_hash"`
	RespondEnabled      bool   `yaml:"respond_to_user_messages"`
	EnableNotifications bool   `yaml:"enable_notifications"`
	ChannelId           int    `yaml:"notification_channel_id"`
	EnableReplays       bool   `yaml:"enable_replays"`
	ChannelIdReplays    int    `yaml:"replays_channel_id"`
	Username            string `yaml:"dlive_username"`
	MaxReplaySizeMb     int    `yaml:"max_replay_size_mb"`
	CachePath           string `yaml:"cache_path"`
	IAEnabled           bool   `yaml:"enable_internet_archive"`
	IAEmail             string `yaml:"internet_archive_email"`
	IAPassword          string `yaml:"internet_archive_password"`
	IAFolderId          string `yaml:"internet_archive_folder"`
	VideoFilename       string `yaml:"video_file_name"`
	ScreenshotFilename  string `yaml:"screenshot_file_name"`
	LogLevel            int    `yaml:"log_level"`
	DebugMode           bool   `yaml:"debug_mode"`
}

var config = Config{
	TelegramBotToken:    "ur_token_goes_here",
	TelegramApiId:       123456,
	TelegramApiHash:     "some_very_long_secret_hash",
	RespondEnabled:      true,
	EnableNotifications: true,
	ChannelId:           123456789,
	EnableReplays:       true,
	ChannelIdReplays:    456789,
	Username:            "nptvpl",
	MaxReplaySizeMb:     2000,
	CachePath:           "./cache",
	IAEnabled:           false,
	IAEmail:             "noreply@my.email",
	IAPassword:          "hackme",
	IAFolderId:          "NPTV-Archive",
	VideoFilename:       "replay.mp4",
	ScreenshotFilename:  "screenshot.jpg",
	LogLevel:            2,
	DebugMode:           false,
}

// Get config somewhere in code
func Get() Config {
	return config
}

// Open file and if exists, read configuration
func Load() bool {
	utils.LogLevel = 10
	file, err := os.ReadFile(fileName)

	// Attempt to create new configuration file
	if err != nil {
		utils.Log("Looks like it's the first run.")
		utils.Check("Creating new config file")

		// Write default config to filesystem
		d, _ := yaml.Marshal(&config)
		err := os.WriteFile(fileName, []byte(d), 0640)
		if err != nil {
			utils.OkFail(false)
			utils.Err("Couldn't create new file. Make sure you have proper permissions to do so.")
			os.Exit(1)
		}

		utils.OkFail(true)
		utils.Log("Before proceeding, please adjust '" + fileName + "' file to your likings.")
		utils.Warn("Make sure to provide valid Telegram creds!")
		utils.Warn("Also, remember to send a message right after bot starts for the first time.")
		os.Exit(0)
	}

	// Try to read existing config
	utils.Check("Loading config file")
	config = Config{}
	err = yaml.Unmarshal([]byte(file), &config)
	ok := err == nil
	utils.OkFail(ok)
	utils.LogLevel = config.LogLevel
	return ok
}
