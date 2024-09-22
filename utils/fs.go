package utils

import (
	"nptw/config"
	"nptw/utils/log"
	"os"
)

func prepareCache() {
	// Prepare filesystem
	log.I("It's time to grab the replay.")
	log.I("Creating cache path")
	err := os.MkdirAll(config.Get().CachePath, 0755)
	if err != nil {
		log.E(err.Error())
	}

}
