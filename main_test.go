package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ewilan-riviere/dlp/pkg/youtube"
)

func TestWebhook(t *testing.T) {
	// Rick Astley - Never Gonna Give You Up (Official Music Video).mp4
	// Rick Astley - Never Gonna Give You Up (Official Music Video).mp3
	// PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC/Rick Astley - Together Forever (Official Video) [Remastered in 4K].mp4
	// PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC/Rick Astley - Never Gonna Give You Up (Official Music Video).mp4
	// yt-dlp -f 'bv*[height=1080]+ba' --download-archive videos.txt  https://www.youtube.com/playlist?list=PLMuc309h9v03vCuZFRO1WAgoPY3CkCUSi -o '%(channel_id)s/%(playlist_id)s/%(id)s.%(ext)s

	youtube.Main(youtube.Params{
		ID:      "dQw4w9WgXcQ",
		Audio:   false,
		FullUrl: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	})

	isExists := fileExists("./Rick Astley - Never Gonna Give You Up (Official Music Video).mp4")
	if !isExists {
		t.Errorf("File not found")
	}

	youtube.Main(youtube.Params{
		ID:      "dQw4w9WgXcQ",
		Audio:   true,
		FullUrl: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	})

	isExists = fileExists("./Rick Astley - Never Gonna Give You Up (Official Music Video).mp3")
	if !isExists {
		t.Errorf("File not found")
	}

	youtube.Main(youtube.Params{
		ID:                 "dQw4w9WgXcQ",
		Audio:              false,
		FullUrl:            "https://www.youtube.com/playlist?list=PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC",
		SaveToDownloadsDir: true,
	})

	// with chapters
	youtube.Main(youtube.Params{
		ID:                 "TLV2IqSIr44",
		Audio:              true,
		FullUrl:            "https://www.youtube.com/watch?v=TLV2IqSIr44",
		SaveToDownloadsDir: true,
		Chapters:           true,
	})

	isExists = fileExists(downloadDirectory() + "/PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC/2-Rick Astley - Together Forever (Official Video) [Remastered in 4K].mp4")
	if !isExists {
		t.Errorf("File not found")
	}

	isExists = fileExists(downloadDirectory() + "/PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC/1-Rick Astley - Never Gonna Give You Up (Official Music Video).mp4")
	if !isExists {
		t.Errorf("File not found")
	}

	isExists = fileExists(downloadDirectory() + "/TLV2IqSIr44/Hazbin Hotel Full Soundtrack - Episodes 1-8 (UPDATED)-NA.mp3")
	if !isExists {
		t.Errorf("File not found")
	}

	youtube.Main(youtube.Params{
		ID:                 "@ewilanriviere2804",
		Audio:              true,
		FullUrl:            "https://www.youtube.com/@ewilanriviere2804",
		SaveToDownloadsDir: true,
	})

	isExists = fileExists(downloadDirectory() + "/@ewilanriviere2804/1-Never Gonna Give You Up - Rick Astley.mp3")
	if !isExists {
		t.Errorf("File not found")
	}

	youtube.Main(youtube.Params{
		ID:                 "@ewilanriviere2804",
		Audio:              false,
		FullUrl:            "https://www.youtube.com/@ewilanriviere2804",
		SaveToDownloadsDir: true,
		Type:               "channel",
	})
	isExists = fileExists(downloadDirectory() + "/@ewilanriviere2804/1-Never Gonna Give You Up - Rick Astley.mp4")
	if !isExists {
		t.Errorf("File not found")
	}
}

func downloadDirectory() string {
	os := runtime.GOOS
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	switch os {
	case "windows":
		return filepath.Join("C:", "Users", username, "Downloads")
	case "darwin":
		return filepath.Join("/Users", username, "Downloads")
	case "linux":
		return filepath.Join("/home", username, "Downloads")
	default:
		return filepath.Join("/home", username, "Downloads")
	}
}

func fileExists(path string) bool {
	fmt.Print("Checking if file exists: " + path)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
