package ffmpeg

import (
	"nptw/utils"
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
	utils.Check("Checking ffmpeg")

	info, err := os.Stat("./bin/ffmpeg")
	ok := !os.IsNotExist(err)
	if ok {
		ok = !info.IsDir()
	}

	utils.OkFail(ok)
	return ok
}

// Install ffmpeg
func Install() bool {
	utils.Check("Installing ffmpeg")
	cmd := exec.Command("bash", "-c", `
		cd bin;
		wget -O ffmpeg.tar.xz https://github.com/BtbN/FFmpeg-Builds/releases/download/latest/ffmpeg-master-latest-linux64-gpl.tar.xz;
		tar -xvf ffmpeg.tar.xz;
		mv ffmpeg*/bin/* .;
		rm -f ffmpeg.tar.xz;
		rm -rf ffmpeg-*;
	`)

	ok := cmd.Run() == nil
	utils.OkFail(ok)
	return ok
}
