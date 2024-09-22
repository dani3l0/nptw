package ia

import (
	"nptw/utils/log"
	"os"
	"os/exec"
)

// YtDlp check and auto-download
func Check() bool {
	exists := Exists()
	if !exists {
		Install()
		exists = Exists()
	}
	return exists
}

// Check if ia is available
func Exists() bool {
	info, err := os.Stat("./bin/ia")
	ok := !os.IsNotExist(err)
	if ok && !info.IsDir() {
		log.I("Internet Archive binary file seems to exist.")
	} else if !ok {
		log.W("Internet Archive binary file does not exist. Downloading now.")
	} else {
		log.E("Path for an Internet Archive binary file is broken.")
		log.E("Try removing whole `bin` directory, setting proper permissions or providing valid path for assets.")
	}
	return ok
}

// Download ia locally
func Install() bool {
	log.I("Installing Interet Archive (ia)")
	cmd := exec.Command("bash", "-c", `
		wget -O ./bin/ia https://archive.org/download/ia-pex/ia
		chmod +x ./bin/ia
	`)

	ok := cmd.Run() == nil

	if ok {
		log.I("Internet Archive installed successfully.")
	} else {
		log.E("Couldn't install Internet Archive properly.")
	}

	return ok
}
