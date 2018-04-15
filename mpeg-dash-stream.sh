#!/bin/bash

CURL=$(which curl)
FFMPEG=$(which ffmpeg)

if [ ! -e "big_buck_bunny_1080p_stereo.ogg" ]
then
	curl https://download.blender.org/peach/bigbuckbunny_movies/big_buck_bunny_1080p_stereo.ogg -o big_buck_bunny_1080p_h264.ogg
fi

OUTPUT="bbb"
DDIR="stream"

$FFMPEG -re -i big_buck_bunny_1080p_stereo.ogg \
	-map 0:0 \
	-pix_fmt yuv420p \
	-codec:v libx264 \
	-b:v 400k -minrate 400k -maxrate 400k -bufsize 800k \
	-r 24 -g 96 -keyint_min 96 -sc_threshold 0 \
	-framerate 24\
	-f segment \
	-segment_format mp4 "${DDIR}/${OUTPUT}_400k_%d.dash" \
	-map 0:1 \
	-codec:a aac \
	-b:a 128k \
	-ar 48000 \
	-f segment \
	-segment_format mp4 "${DDIR}/${OUTPUT}_128k_%d.dash"

