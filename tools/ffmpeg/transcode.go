package ffmpeg

import (
	"fmt"
	"nptw/config"
	"nptw/utils"
	"os/exec"
	"path"
	"strings"
)

// Install ffmpeg
func Transcode(bitrate float64) bool {
	var cmd string
	utils.Log("Preparing ffmpeg for video transcoding")

	// Some variables for ffmpeg
	ffbin := "./bin/ffmpeg"
	vid_src := path.Join(config.Get().CachePath, "replay.mp4")
	vid_tgt := path.Join(config.Get().CachePath, "output.mp4")

	// Prepare ffmpeg commands
	kbps := int(bitrate * 1000)
	cmd_oncpu := fmt.Sprintf(
		"%s -y -i %s -c:v h264 -preset veryfast -b:v %dk %s",
		ffbin, vid_src, kbps, vid_tgt,
	)
	cmd_libva := fmt.Sprintf(
		"%s -y -hwaccel vaapi -hwaccel_device %s -hwaccel_output_format vaapi -i %s -c:v h264_vaapi -b:v %dk %s",
		ffbin, config.Get().FfmpegHwAccelDevice, vid_src, kbps, vid_tgt,
	)
	cmd_onqsv := fmt.Sprintf(
		"%s -y -hwaccel qsv -qsv_device %s -hwaccel_output_format qsv -i %s -c:v h264_qsv -b:v %dk %s",
		ffbin, config.Get().FfmpegHwAccelDevice, vid_src, kbps, vid_tgt,
	)

	// Select proper hwaccel mode
	if config.Get().FfmpegHwAccel && config.Get().UseQSV {
		cmd = cmd_onqsv
	} else if config.Get().FfmpegHwAccel {
		cmd = cmd_libva
	} else {
		cmd = cmd_oncpu
	}

	// Actually, run transcoding
	e := exec.Command(cmd)
	res, err := e.CombinedOutput()
	if err == nil {
		utils.Log("FFmpeg transcoding successful")
	} else {
		utils.Err("FFmpeg transcoding failed")
	}

	utils.Log("FFmpeg output:")
	for _, v := range strings.Split(string(res), "\n") {
		utils.Log(v)
	}

	return err == nil
}
