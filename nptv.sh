#!/bin/bash

url="https://live.prd.dlive.tv/hls/live/nptvpl.m3u8"
timeout=10

# Start an infinite loop
while true; do
    # Get the current hour without leading zero
    current_hour=$(date "+%-H")

    # Check if the current time is not between 01:00 and 10:00
    if (( current_hour < 1 || current_hour >= 10 )); then
        if ffprobe -v error -show_entries stream=codec_type -of default=nokey=1:noprint_wrappers=1 "$url" | grep -q "video"; then
            # Stream is live
            echo "Stream is live"





miniaturka=$(yt-dlp -j https://dlive.tv/nptvpl | jq -r '.thumbnails[0].url')



			
		if [ -z "$miniaturka" ]; then
       
        links=("https://przykladoweminiaturki.com/ryciny/dd.jpg" "https://przykladoweminiaturki.com/ryciny/dd2.jpg" "https://przykladoweminiaturki.com/ryciny/dd3.jpg" "https://przykladoweminiaturki.com/ryciny/dd4.jpg" "https://przykladoweminiaturki.com/ryciny/dd5.jpg")
        random_index=$(($RANDOM % ${#links[@]}))
        miniaturka=${links[$random_index]}

else
        wget $miniaturka -O /tmp/obrazek.jpg
        cd /tmp/
        image_path="obrazek.jpg"
        brightness_threshold=0.25
        get_brightness() {
            brightness=$(convert "$1" -colorspace gray -format "%[fx:mean]" info:)
            echo "$brightness"
}
brightness=$(get_brightness "$image_path")
if (( $(bc <<< "$brightness >= $brightness_threshold") )); then
    echo "jasny"
else
    echo "ciemny"
            links=("https://przykladoweminiaturki.com/ryciny/dd.jpg" "https://przykladoweminiaturki.com/ryciny/dd2.jpg" "https://przykladoweminiaturki.com/ryciny/dd3.jpg" "https://przykladoweminiaturki.com/ryciny/dd4.jpg" "https://przykladoweminiaturki.com/ryciny/dd5.jpg" "https://przykladoweminiaturki.com/ryciny/dd6.jpg" "https://przykladoweminiaturki.com/ryciny/dd7.jpg" "https://przykladoweminiaturki.com/ryciny/dd8.jpg" "https://przykladoweminiaturki.com/ryciny/dd9.jpg" "https://przykladoweminiaturki.com/ryciny/dd10.jpg")
        random_index=$(($RANDOM % ${#links[@]}))
        miniaturka=${links[$random_index]}
fi
		fi




		if [ -z "$title" ]; then

		title=$(yt-dlp -j https://dlive.tv/nptvpl | jq -r '.fulltitle')
		title="${title%%RodacyKamraci, Wojciech Olszański, Marcin Osadowski*}"
		fi

		if [ -z "$title" ]; then
		title=""
		else
		title="🐺*$title*🦎"
		lowq=$(yt-dlp -j https://dlive.tv/nptvpl | jq -r '.formats[] | select(.format_id == "360p") | .url')
		lowq="[Link do żywca dla słabszych łącz]($lowq)"
		fi







curl -s -X POST "https://api.telegram.org/botXXXX/sendPhoto" -d "chat_id=@zywce" -d 'caption=🇵🇱Rozpoczął się *żywiec*!🇵🇱%0A %0A '"$title"'%0A %0A➡Link do dlive: https://dlive.tv/nptvpl %0A' -d photo="$miniaturka" -d "parse_mode=Markdown"



rm /tmp/obrazek.jpg

#zeruj zmienne
title=""
miniaturka=""

            sleep 21600

        else
            # Stream is not live
            echo "Stream is not live"
            # Wait for 5 minutes before checking again
            sleep 300
        fi
    else
        # Wait for 15 minutes before checking again
        sleep 900
    fi
done


#archiwalne linki/powtórki
#curl 'https://graphigo.prd.dlive.tv/'  --data-raw '{"operationName":"LivestreamProfileReplay","variables":{"displayname":"NPTVPL","first":20},"extensions":{"persistedQuery":{"version":1,"sha256Hash":"913ee036eb38d230cdc0d46c8d3a0b016227b3ee2bf3f6b2f501a0a238d43e9a"}}}' | jq -r '.data.userByDisplayName.pastBroadcasts.list[].permlink'
