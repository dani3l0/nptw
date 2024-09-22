package ytdlp

import (
	"fmt"
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"os/exec"
	"path"
	"strings"
)

// Get info about current stream
func GetInfo(url string) (string, error) {
	log.I("yt-dlp is getting info about " + url)
	cmd := exec.Command("./bin/yt-dlp", "-J", url)
	output, err := cmd.Output()
	if err != nil {
		log.E("yt-dlp couldn't get info about current stream!")
	}
	return string(output[:]), err
}

// Fetch raw stream file in chunks so we can pass it to pipe
func GetYtDlpFfmpegCmd(url string) []string {
	hwaccelDevice := config.Get().FfmpegHwAccelDevice
	quality := config.Get().FfmpegReplayWidthPixels
	hwaccelType := config.Get().FfmpegHwAccelType

	// CPU: no hardware acceleration
	hwaccel := []string{
		"-preset", "veryfast",
		"-c:v", "h264",
	}

	if hwaccelType == "qsv" {
		// Intel QuickSyncVideo on iGPU
		hwaccel = []string{
			"-hwaccel", "qsv",
			"-qsv_device", hwaccelDevice,
			"-hwaccel_output_format", "qsv",
			"-c:v", "h264_qsv",
		}

	} else if hwaccelType == "vaapi" {
		// VAAPI, universal for AMD, Intel and possibly NVIDIA
		hwaccel = []string{
			"-hwaccel", "vaapi",
			"-hwaccel_device", hwaccelDevice,
			"-hwaccel_output_format", "vaapi",
			"-c:v", "h264_vaapi",
		}

	} else if hwaccelType != "cpu" {
		// Just a warn when config has unsupported value set
		log.W("ffmpeg_hwaccel_type was provided with invalid value. Supported ones are: qsv, vaapi, cpu. Falling back to cpu.")
	}

	var cmd []string

	// Build target ffmpeg command
	cmd = append(cmd, "./bin/ffmpeg")                                            // path to ffmpeg
	cmd = append(cmd, hwaccel...)                                                // apply hardware acceleration parameters
	cmd = append(cmd, "-i", fmt.Sprintf("$(./bin/yt-dlp '%s' -f best -g)", url)) // get raw stream to convert it on-the-fly
	cmd = append(cmd, fmt.Sprintf("-filter:v scale=%d:-1", quality))             // resize video to width specified in config file
	cmd = append(cmd, "-b:a", "128k")                                            // set audio bitrate to 128kbps to save space
	cmd = append(cmd, path.Join(config.Get().CachePath, globals.VideoFilename))  // where to save converted replay file

	log.I("Generated hwaccel cmd:`", strings.Join(cmd, " "), "`")

	return cmd
}
