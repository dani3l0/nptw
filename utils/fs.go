package utils

import (
	"nptw/config"
	"nptw/utils/log"
	"os"
)

func PrepareCache() {
	// Prepare filesystem
	log.I("Creating cache path")
	err := os.MkdirAll(config.Get().CachePath, 0755)
	if err != nil {
		log.E(err.Error())
	}
}

func CleanCache() {
	// Nuke temporary folder
	err := os.RemoveAll(config.Get().CachePath)
	if err != nil {
		log.E("Cleaning up cache failed: ", err.Error())
	}
}
