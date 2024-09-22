package ffmpeg

import (
	"nptw/utils/log"
	"os"
	"os/exec"
)

// Ffmpeg check and auto-download
func Check() bool {
	exists := Exists()
	if !exists {
		Install()
		exists = Exists()
	}
	return exists
}

// Check if ffmpeg is installed
func Exists() bool {
	info, err := os.Stat("./bin/ffmpeg")
	ok := !os.IsNotExist(err)
	if ok && !info.IsDir() {
		log.I("ffmpeg binary file seems to exist.")
	} else if !ok {
		log.W("ffmpeg binary file does not exist. Downloading now.")
	} else {
		log.E("Path for an ffmpeg binary file is broken.")
		log.E("Try removing whole `bin` directory, setting proper permissions or providing valid path for assets.")
	}
	return ok
}

// Install ffmpeg
func Install() bool {
	log.I("Installing ffmpeg")
	cmd := exec.Command("bash", "-c", `
		cd bin;
		wget -O ffmpeg.tar.xz https://github.com/BtbN/FFmpeg-Builds/releases/download/latest/ffmpeg-master-latest-linux64-gpl.tar.xz;
		tar -xvf ffmpeg.tar.xz;
		mv ffmpeg*/bin/* .;
		rm -f ffmpeg.tar.xz;
		rm -rf ffmpeg-*;
	`)

	ok := cmd.Run() == nil

	if ok {
		log.I("ffmpeg installed successfully.")
	} else {
		log.E("Couldn't install ffmpeg properly.")
	}

	return ok
}
