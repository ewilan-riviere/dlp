// Package youtube-dl is a CLI to use yt-dlp (https://github.com/yt-dlp/yt-dlp) easily.
//
// `get` command download a video, a playlist or a channel. Argument can be a URL or an YouTube ID, it will be detected automatically.
// With URL, it possible to use another website than YouTube, like DailyMotion or PeerTube for example, ID only accept YouTube ID (video, playlist or channel).
//
// - Videos will be downloaded in `Downloads` folder
// - You can execute command without argument, it will ask you to enter an URL or an ID.
//
// Options:
//
// - `-a` or `--audio`: download only audio
//
// Examples:
//
//	dlp get <URL_OR_ID>
//	dlp get <URL_OR_ID> -a
//
// You can skip argument and enter URL or ID after command execution:
//
//	dlp get
//
// You can also use it with another website than YouTube:
//
//	dlp get https://www.dailymotion.com/video/x3lnp8j
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ewilan-riviere/dlp/pkg/youtube"
	"github.com/spf13/cobra"
)

func main() {
	var cmdGet = &cobra.Command{
		Use:   "get [id]",
		Short: "Download video from URL or Youtube from ID",
		Long:  `Download video from URL or Youtube from ID, it can be a video, a playlist or a channel`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			audio, _ := cmd.Flags().GetBool("audio")
			id := ""

			if len(args) > 0 {
				id = args[0]
			}

			if id == "" {
				fmt.Println("")
				fmt.Println("Please enter a video URL or a Youtube ID")
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("URL or ID: ")
				text, _ := reader.ReadString('\n')
				id = strings.Trim(text, "\n")
			}

			url := ""
			isUrl := false
			if strings.Contains(id, "http") {
				url = id
				isUrl = true
			}

			if !isUrl {
				url = youtubeUrl(id)
			}

			youtube.Main(youtube.Params{
				Audio:   audio,
				FullUrl: url,
			})
		},
	}

	cmdGet.Flags().BoolP("audio", "a", false, "To convert downloaded video to audio (works with video, playlist, channel)")
	var rootCmd = &cobra.Command{Use: "dlp"}
	rootCmd.AddCommand(cmdGet)
	rootCmd.Execute()
}

func youtubeUrl(id string) string {
	const YoutubeBaseUrl = "https://www.youtube.com"
	const YoutubeTypeVideo = "watch?v="
	const YoutubeTypePlaylist = "playlist?list="
	const YoutubeTypeChannel = "/"

	if strings.Contains(id, "@") {
		return YoutubeBaseUrl + "/" + YoutubeTypeChannel + id
	}

	length := len(id)
	if length > 15 {
		return YoutubeBaseUrl + "/" + YoutubeTypePlaylist + id
	}

	return YoutubeBaseUrl + "/" + YoutubeTypeVideo + id
}
