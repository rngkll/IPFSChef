#!/usr/bin/env bash

INPUT="big_buck_bunny_1080p_stereo.ogg"

# Video encoding from live
VP9_LIVE_PARAMS="-speed 6 -tile-columns 4 -frame-parallel 1 -threads 8 -static-thresh 0 -max-intra-rate 300 -deadline realtime -lag-in-frames 0 -error-resilient 1"

ffmpeg \
	-f v4l2 -input_format mjpeg -r 30 -s 1280x720 -i /dev/video0 \
	-f alsa -ar 44100 -ac 2 -i hw:2 \
	-map 0:0 \
	-pix_fmt yuv420p \
	-c:v libvpx-vp9 \
	-s 1280x720 -keyint_min 60 -g 60 ${VP9_LIVE_PARAMS} \
	-b:v 3000k \
	-f webm_chunk \
	-header "/var/www/webm_live/glass_360.hdr" \
	-chunk_start_index 1 \
	webm_live/glass_360_%d.chk \
	-map 1:0 \
	-c:a libvorbis \
	-b:a 128k -ar 44100 \
	-f webm_chunk \
	-audio_chunk_duration 2000 \
	-header webm_live/glass_171.hdr \
	-chunk_start_index 1 \
	webm_live/glass_171_%d.chk

# Reference
# http://wiki.webmproject.org/adaptive-streaming/instructions-to-playback-adaptive-webm-using-dash
