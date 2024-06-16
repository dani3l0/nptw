package tools

import (
	"nptw/utils"
	"os"
	"os/exec"
)

// Ffmpeg check and auto-download
func FfmpegCheck() bool {
	exists := FfmpegExists()
	if !exists {
		FfmpegInstall()
		exists = FfmpegExists()
	}
	return exists
}

// Check if ffmpeg is installed
func FfmpegExists() bool {
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
func FfmpegInstall() bool {
	utils.Check("Installing ffmpeg ...")
	cmd := exec.Command("bash", "-c", `
		cd bin;
		wget -O ffmpeg.tar.xz https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz;
		tar -xvf ffmpeg.tar.xz;
		mv ffmpeg*/* .;
		rm ffmpeg.tar.xz;
		rm -r ffmpeg-*-static;
	`)

	ok := cmd.Run() == nil
	utils.OkFail(ok)
	return ok
}
