package ytdlp

import (
	"nptw/utils/log"
	"os/exec"
	"strings"
)

// Fetch raw stream file in chunks so we can pass it to pipe
func GetYtDlpFfmpegCmd(url string) []string {
	// hwaccelDevice := config.Get().FfmpegHwAccelDevice
	// quality := config.Get().FfmpegReplayWidthPixels
	// hwaccelType := config.Get().FfmpegHwAccelType

	// CPU: no hardware acceleration
	// hwaccel := []string{
	// 	"-preset", "veryfast",
	// 	"-c:v", "h264",
	// }

	// if hwaccelType == "qsv" {
	// 	// Intel QuickSyncVideo on iGPU
	// 	hwaccel = []string{
	// 		"-hwaccel", "qsv",
	// 		"-qsv_device", hwaccelDevice,
	// 		"-hwaccel_output_format", "qsv",
	// 		"-c:v", "h264_qsv",
	// 	}

	// } else if hwaccelType == "vaapi" {
	// 	// VAAPI, universal for AMD, Intel and possibly NVIDIA
	// 	hwaccel = []string{
	// 		"-hwaccel", "vaapi",
	// 		"-hwaccel_device", hwaccelDevice,
	// 		"-hwaccel_output_format", "vaapi",
	// 		"-c:v", "h264_vaapi",
	// 	}

	// } else if hwaccelType != "cpu" {
	// 	// Just a warn when config has unsupported value set
	// 	log.W("ffmpeg_hwaccel_type was provided with invalid value. Supported ones are: qsv, vaapi, cpu. Falling back to cpu.")
	// }

	var cmd, cmd1, cmd2 []string

	cmd1 = append(cmd1, "-f", "best", "-o", "-", "https://dlive.tv/p/nptvpl+OEPKp3eIg")
	cmd2 = append(cmd2, "-y", "-i", "-")
	cmd2 = append(cmd2, "-vf", "scale=1280:-2")
	cmd2 = append(cmd2, "-b:v 512k")
	cmd2 = append(cmd2, "-c:v libx264")
	cmd2 = append(cmd2, "-preset veryfast")
	cmd2 = append(cmd2, "-b:a 128k")
	cmd2 = append(cmd2, "output.mp4")

	dlp := exec.Command("./bin/yt-dlp", cmd1...)
	mpg := exec.Command("./bin/ffmpeg", cmd2...)
	log.I("Generated hwaccel cmd:")
	log.I(strings.Join(cmd, " "))

	dlp.Stdin, _ = mpg.StdoutPipe()
	// dlp.Stdout = os.Stdout
	_ = dlp.Start()
	_ = mpg.Run()
	_ = dlp.Wait()

	return cmd
}
