package main

import (
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"github.com/zencoder/go-dash/mpd"
)

var sh *shell.Shell

func main() {

	m := mpd.NewMPD(mpd.DASH_PROFILE_LIVE, "PT6M16S", "PT1.97S")

	audioAS, _ := m.AddNewAdaptationSetAudio("audio", mpd.DASH_MIME_TYPE_AUDIO_MP4, true, 1, "und")

	audioAS.SetNewSegmentTemplate(1968, "$RepresentationID$/audio/en/init.mp4", "$RepresentationID$/audio/en/seg-$Number$.m4f", 0, 1000)
	audioAS.AddNewRepresentationAudio(44100, 67095, "mp4a.40.2", "800")

	videoAS, _ := m.AddNewAdaptationSetVideo("video", mpd.DASH_MIME_TYPE_VIDEO_MP4, "progressive", true, 1)

	videoAS.SetNewSegmentTemplate(1968, "$RepresentationID$/video/1/init.mp4", "$RepresentationID$/video/1/seg-$Number$.m4f", 0, 1000)
	videoAS.AddNewRepresentationVideo(1518664, "avc1.4d401f", "800", "30000/1001", 960, 540)
	videoAS.AddNewRepresentationVideo(1911775, "avc1.4d401f", "1000", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2295158, "avc1.4d401f", "1200", "30000/1001", 1024, 576)
	videoAS.AddNewRepresentationVideo(2780732, "avc1.4d401f", "1500", "30000/1001", 1280, 720)

	mpdStr, _ := m.WriteToString()
	fmt.Println(mpdStr)

	sh = shell.NewShell("localhost:5001")
	sh.PubSubPublish("IpfsChefChannel", mpdStr)
}
