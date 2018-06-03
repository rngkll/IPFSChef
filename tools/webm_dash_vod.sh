#!/usr/bin/env bash

INPUT="Sintel.2010.1080p.mkv"

# Video encoding from raw
VP9_DASH_PARAMS="-tile-columns 4 -frame-parallel 1"

# ffmpeg -i ${INPUT} -c:v libvpx-vp9 -s 160x90 -b:v 250k -keyint_min 150 -g 150 ${VP9_DASH_PARAMS} -an -f webm -dash 1 webm_vod/video_160x90_250k.webm
# ffmpeg -i ${INPUT} -c:v libvpx-vp9 -s 320x180 -b:v 500k -keyint_min 150 -g 150 ${VP9_DASH_PARAMS} -an -f webm -dash 1 webm_vod/video_320x180_500k.webm
# # ffmpeg -i ${INPUT} -c:v libvpx-vp9 -s 640x360 -b:v 750k -keyint_min 150 -g 150 ${VP9_DASH_PARAMS} -an -f webm -dash 1 webm_vod/video_640x360_750k.webm
# # ffmpeg -i ${INPUT} -c:v libvpx-vp9 -s 640x360 -b:v 1000k -keyint_min 150 -g 150 ${VP9_DASH_PARAMS} -an -f webm -dash 1 webm_vod/video_640x360_1000k.webm
# # ffmpeg -i ${INPUT} -c:v libvpx-vp9 -s 1280x720 -b:v 1500k -keyint_min 150 -g 150 ${VP9_DASH_PARAMS} -an -f webm -dash 1 webm_vod/video_1280x720_500k.webm
# 
# # Audio encoding from raw
# ffmpeg -i ${INPUT} -c:a libvorbis -b:a 128k -vn -f webm -dash 1 webm_vod/audio_128k.webm

# Create manifest
ffmpeg \
	-f webm_dash_manifest -i webm_vod/video_160x90_250k.webm \
	-f webm_dash_manifest -i webm_vod/video_320x180_500k.webm \
	-f webm_dash_manifest -i webm_vod/audio_128k.webm \
	-c copy -map 0:0 -map 1:0 -map 2:0 \
	-f webm_dash_manifest \
	-adaptation_sets "id=0,webm_vods=0,1 id=1,webm_vods=2" \
	webm_vod/manifest.mpd

# Reference
# http://wiki.webmproject.org/adaptive-webm_voding/instructions-to-playback-adaptive-webm-using-dash
