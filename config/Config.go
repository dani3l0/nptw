package config

import (
	"nptw/utils"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramBotToken string
	TelegramApiId    int
	TelegramApiHash  string
	BotOwnerId       int
	ChannelId        int
	CachePath        string
}

var config = Config{
	CachePath: "./cache",
}

// Get config somewhere in code
func Get() Config {
	return config
}

// Open file and if exists, read configuration
func Load() bool {
	file, err := os.ReadFile("config.yml")

	// Attempt to create new configuration file
	if err != nil {
		utils.Check("Config file not found, creating one")

		d, _ := yaml.Marshal(&config)
		err := os.WriteFile("config.yaml", []byte(d), 0640)
		if err != nil {
			utils.OkFail(false)
			utils.Log("Couldn't create new file. Make sure you have proper permissions to do so.")
			return false
		}

		utils.OkFail(true)
		utils.Log("Before proceeding, please adjust 'config.yaml' file to your likings.")
		utils.Log("Make sure to provide valid Telegram creds!")
		return false
	}

	// Try to read existing config
	err = yaml.Unmarshal(file, &config)
	ok := err != nil
	utils.Check("Loading config file")
	utils.OkFail(ok)
	return ok
}
