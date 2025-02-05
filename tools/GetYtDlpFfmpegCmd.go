package tools

import (
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"path"
	"runtime"
	"strconv"
)

// Fetch raw stream file in chunks so we can pass it to pipe
func GetYtDlpFfmpegCmd(url string, videoBitrate int) []string {
	var cmd []string
	hwaccelDevice := config.Get().FfmpegHwAccelDevice
	hwaccelType := config.Get().FfmpegHwAccelType
	tmpVidPath := path.Join(config.Get().CachePath, globals.VideoFilename)
	useHevc := config.Get().FfmpegHevc
	codec := "h264"
	if useHevc {
		codec = "hevc"
	}
	log.I("ffmpeg: Encoding with ", codec)

	// Ffmpeg CPU threads to be used
	threads := runtime.NumCPU()
	if threads >= 4 {
		threads /= 2
	}
	if config.Get().FfmpegThreads > 0 {
		threads = config.Get().FfmpegThreads
	}
	log.I("ffmpeg: used threads: " + strconv.Itoa(threads))

	if hwaccelType == "vaapi" {
		// VAAPI, universal for AMD, Intel and possibly NVIDIA
		cmd = []string{
			"./bin/ffmpeg", "-y",
			"-hwaccel_device", hwaccelDevice,
			"-hwaccel", "vaapi",
			"-hwaccel_output_format", "vaapi",
			"-i", "$(./bin/yt-dlp '" + url + "' -g)",
			"-threads", strconv.Itoa(threads),
			"-b:a", strconv.Itoa(globals.AudioBitrate) + "k",
			"-c:v", codec + "_vaapi",
			"-b:v", strconv.Itoa(videoBitrate) + "k",
			tmpVidPath,
		}

	} else if hwaccelType == "qsv" {
		// Intel QuickSyncVideo on iGPU
		cmd = []string{
			"./bin/ffmpeg", "-y",
			"-qsv_device", hwaccelDevice,
			"-hwaccel", "qsv",
			"-hwaccel_output_format", "qsv",
			"-c:v", "h264_qsv",
			"-i", "$(./bin/yt-dlp '" + url + "' -g)",
			"-threads", strconv.Itoa(threads),
			"-b:a", strconv.Itoa(globals.AudioBitrate) + "k",
			"-c:v", codec + "_qsv",
			"-b:v", strconv.Itoa(videoBitrate) + "k",
			tmpVidPath,
		}

	} else if hwaccelType == "cuda" {
		// NVIDIA hardware acceleration using dGPU video engines
		cmd = []string{
			"./bin/ffmpeg", "-y",
			"-hwaccel_device", hwaccelDevice,
			"-hwaccel", "cuda",
			"-hwaccel_output_format", "cuda",
			"-i", "$(./bin/yt-dlp '" + url + "' -g)",
			"-threads", strconv.Itoa(threads),
			"-b:a", strconv.Itoa(globals.AudioBitrate) + "k",
			"-c:v", codec + "_nvenc",
			"-b:v", strconv.Itoa(videoBitrate) + "k",
			tmpVidPath,
		}

	} else {
		// CPU, uses a LOT OF POWER and generates SO MUCH HEAT
		if hwaccelDevice != "cpu" {
			log.W("ffmpeg_hwaccel_type was provided with invalid value `", hwaccelType, "`. Supported ones are: qsv, vaapi, cuda, cpu. Falling back to cpu.")
		}
		if config.Get().FfmpegHevc {
			log.W("ffmpeg: ignoring `hevc` flag as it is too heavy for CPU! Falling back to h264.")
		}
		cmd = []string{
			"./bin/ffmpeg", "-y",
			"-i", "$(./bin/yt-dlp '" + url + "' -g)",
			"-threads", strconv.Itoa(threads),
			"-b:a", strconv.Itoa(globals.AudioBitrate) + "k",
			"-b:v", strconv.Itoa(videoBitrate) + "k",
			tmpVidPath,
		}
	}

	return cmd
}
