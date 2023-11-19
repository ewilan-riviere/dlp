package main

import (
	"testing"

	"github.com/ewilan-riviere/dlp/pkg/youtube"
)

func TestWebhook(t *testing.T) {
	youtube.Main(youtube.Params{
		ID:      "dQw4w9WgXcQ",
		Audio:   false,
		FullUrl: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	})

	youtube.Main(youtube.Params{
		ID:      "dQw4w9WgXcQ",
		Audio:   true,
		FullUrl: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	})

	youtube.Main(youtube.Params{
		ID:      "dQw4w9WgXcQ",
		Audio:   false,
		FullUrl: "https://www.youtube.com/playlist?list=PLomb1f2d9BZrQc82QHJEDYgnPXHfMDjlC",
	})
}
