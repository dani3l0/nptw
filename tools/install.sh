#!/bin/bash

# Prepare dir
mkdir -p bin
cd bin

# yt dlp
wget -O yt-dlp https://github.com/yt-dlp/yt-dlp/releases/download/2024.05.27/yt-dlp_linux
chmod +x yt-dlp

# ffmpeg
wget -O ffmpeg.tar.xz https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
tar -xvf ffmpeg.tar.xz
mv ffmpeg*/* .
rm ffmpeg.tar.xz
rm -r ffmpeg-*-static
