package config

import (
	"nptw/utils"
	"os"

	"gopkg.in/yaml.v3"
)

var fileName = "config.yaml"

type Config struct {
	TelegramBotToken    string
	TelegramApiId       int
	TelegramApiHash     string
	ChannelId           int
	CachePath           string
	Username            string
	MaxReplaySizeMb     int
	MaxUploadSizeMb     int
	FfmpegHwAccel       bool
	FfmpegHwAccelDevice string
	UseQSV              bool
	LogLevel            int
	IAEmail             string
	IAPassword          string
	IAFolderId          string
}

var config = Config{}

// Get config somewhere in code
func Get() Config {
	return config
}

// Open file and if exists, read configuration
func Load() bool {
	file, err := os.ReadFile(fileName)

	// Attempt to create new configuration file
	if err != nil {
		utils.Log("Looks like it's the first run.")
		utils.Check("Creating new config file")

		// Default config
		config = Config{
			TelegramBotToken:    "ur_token_goes_here",
			TelegramApiId:       123456,
			TelegramApiHash:     "some_very_long_secret_hash",
			ChannelId:           123456789,
			CachePath:           "./cache",
			Username:            "nptvpl",
			MaxReplaySizeMb:     2000,
			MaxUploadSizeMb:     1000,
			FfmpegHwAccel:       false,
			FfmpegHwAccelDevice: "/dev/dri/renderD128",
			UseQSV:              true,
			LogLevel:            2,
			IAEmail:             "noreply@my.email",
			IAPassword:          "hackme",
			IAFolderId:          "NPTV-Archive",
		}

		// Write default config to filesystem
		d, _ := yaml.Marshal(&config)
		err := os.WriteFile(fileName, []byte(d), 0640)
		if err != nil {
			utils.OkFail(false)
			utils.Log("Couldn't create new file. Make sure you have proper permissions to do so.")
			os.Exit(1)
		}

		utils.OkFail(true)
		utils.Log("Before proceeding, please adjust '" + fileName + "' file to your likings.")
		utils.Log("Make sure to provide valid Telegram creds!")
		os.Exit(0)
	}

	// Try to read existing config
	utils.Check("Loading config file")
	err = yaml.Unmarshal([]byte(file), &config)
	ok := err == nil
	utils.OkFail(ok)
	utils.LogLevel = config.LogLevel
	return ok
}
