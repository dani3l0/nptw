package config

import (
	"nptw/utils"
	"os"

	"gopkg.in/yaml.v3"
)

var fileName = "config.yaml"

type Config struct {
	TelegramBotToken string
	TelegramApiId    int
	TelegramApiHash  string
	ChannelId        int
	CachePath        string
	Username         string
}

var config = Config{
	CachePath: "./cache",
	Username:  "nptvpl",
}

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
	return ok
}
