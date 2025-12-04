package tools

import (
	"fmt"
	"nptw/config"
	"nptw/config/globals"
	"nptw/utils/log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// Fetch raw stream file in chunks so we can pass it to pipe
func GetYtDlpFfmpegCmd(url string, videoBitrate int) []string {
	var cmd []string
	tmpVidPath := path.Join(config.Get().CachePath, globals.VideoFilename)

	// Ffmpeg CPU threads to be used
	cpus := runtime.NumCPU()
	threads := cpus
	if threads >= 4 {
		threads /= 2
	}
	if config.Get().FfmpegThreads > 0 {
		threads = config.Get().FfmpegThreads
		if threads > cpus {
			threads = cpus
		}
	}
	log.I("ffmpeg: Used threads: " + strconv.Itoa(threads))

	// Generic starter command
	cachePath := config.Get().CachePath
	cwd, _ := os.Getwd()
	ytdlp := path.Join(cwd, "bin/yt-dlp")
	ffmpeg := path.Join(cwd, "bin/ffmpeg")
	ythreads := 1
	if config.Get().YtDlpThreads > cpus {
		ythreads = cpus
	} else if config.Get().YtDlpThreads < 1 {
		ythreads = 1
	}
	cmd = []string{
		"cd", cachePath, ";",
		ytdlp, "-o", "-", fmt.Sprintf(`'%s'`, url),
		"--concurrent-fragments", strconv.Itoa(ythreads),
		"--cache-dir", cachePath, "--paths", cachePath,
		"|",
		ffmpeg, "-y", "-i", "pipe:0", "-fflags", "+discardcorrupt",
		"-threads", strconv.Itoa(threads),
		"-b:a", strconv.Itoa(globals.AudioBitrate) + "k",
		"-b:v", strconv.Itoa(videoBitrate) + "k",
	}

	// Encoder & hwaccel selection
	var codec string
	acceltype := config.Get().FfmpegHwAccelType
	switch acceltype {
	case "cuda":
		codec = encoder("h264_nvenc", "hevc_nvenc", "av1_nvenc")
	case "vaapi":
		codec = encoder("h264_vaapi", "hevc_vaapi", "av1_vaapi")
	case "qsv":
		codec = encoder("h264_qsv", "hevc_qsv", "av1_qsv")
	default:
		codec = encoder("libx264", "libx265", "libaom-av1")
		log.W("ffmpeg hardware acceleration disabled. Expect high CPU usage and long transcodes. Also, this will make your room warmer...")
	}
	cmd = append(cmd, "-c:v", codec)
	cmd = append(cmd, tmpVidPath)

	return cmd
}

// Selects encoder format according to config
func encoder(x264 string, x265 string, av1 string) string {
	f := strings.ToLower(config.Get().FfmpegFormat)
	switch f {
	case "av1":
		return av1
	case "h265", "hevc":
		return x265
	default:
		return x264
	}
}
