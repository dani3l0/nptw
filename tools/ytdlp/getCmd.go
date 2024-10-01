package ytdlp

import (
	"fmt"
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"path"
	"runtime"
	"strconv"
)

// Fetch raw stream file in chunks so we can pass it to pipe
func GetYtDlpFfmpegCmd(url string, videoBitrate int) []string {
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
			"-qsv_device", hwaccelDevice,
			"-hwaccel_output_format", "qsv",
			"-c:v", "h264_qsv",
		}

	} else if hwaccelType == "vaapi" {
		// VAAPI, universal for AMD, Intel and possibly NVIDIA
		hwaccel = []string{
			"-hwaccel_device", hwaccelDevice,
			"-hwaccel_output_format", "vaapi",
			"-c:v", "h264_vaapi",
		}

	} else if hwaccelType != "cpu" {
		// Just a warn when config has unsupported value set
		log.W("ffmpeg_hwaccel_type was provided with invalid value `", hwaccelType, "`. Supported ones are: qsv, vaapi, cpu. Falling back to cpu.")
	}

	// Ffmpeg CPU threads to be used
	threads := runtime.NumCPU()
	if threads >= 4 {
		threads /= 2
	} else if config.Get().FfmpegThreads > 0 {
		threads = config.Get().FfmpegThreads
	}

	// Build magic command
	var cmd []string
	cmd = append(cmd, "./bin/ffmpeg", "-y")
	cmd = append(cmd, hwaccel...)
	cmd = append(cmd, "-i", fmt.Sprintf("$(./bin/yt-dlp -f best %s -g)", url))
	cmd = append(cmd, fmt.Sprintf("-vf scale=%d:-2", quality))
	cmd = append(cmd, "-threads", strconv.Itoa(threads))
	cmd = append(cmd, "-b:v", strconv.Itoa(videoBitrate)+"k")
	cmd = append(cmd, "-b:a", strconv.Itoa(globals.AudioBitrate)+"k")
	cmd = append(cmd, path.Join(config.Get().CachePath, globals.VideoFilename))

	return cmd
}
